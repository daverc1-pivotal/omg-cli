package bosh

import (
	"fmt"
	"time"

	"github.com/enaml-ops/enaml"
	"github.com/enaml-ops/enaml/enamlbosh"
	"github.com/enaml-ops/pluginlib/cloudconfig"
	"github.com/enaml-ops/pluginlib/pcli"
	"github.com/enaml-ops/pluginlib/product"
	"github.com/xchapter7x/lo"
	"gopkg.in/urfave/cli.v2"
)

// Reset all custom styles
const RESET = "\033[0m"

// Return curor to start of line and clean it
const RESET_LINE = "\r\033[K"

var UIPrint = fmt.Println

var UIPrintStatus = func(i ...interface{}) (int, error) {
	fmt.Print(RESET_LINE)
	return fmt.Printf(i[0].(string)+getProcessing(), i[1:]...)
}

var boshclient *enamlbosh.Client

// getBoshClient implements lazy instantiation for this
// package's enamlbosh Client
func getBoshClient(c *cli.Context) *enamlbosh.Client {
	if boshclient == nil {
		boshUser := c.String("bosh-user")
		boshPass := c.String("bosh-pass")
		boshURL := c.String("bosh-url")
		boshPort := c.Int("bosh-port")
		skipSSLVerify := c.Bool("ssl-ignore")

		var err error
		boshclient, err = enamlbosh.NewClient(boshUser, boshPass, boshURL, boshPort, skipSSLVerify)

		if err != nil {
			lo.G.Panic("Couldn't create bosh client:", err)
		}
	}
	return boshclient
}

func GetAuthFlags() []pcli.Flag {
	return []pcli.Flag{
		pcli.CreateStringFlag("bosh-url", "the url or ip of your bosh director", "https://mybosh.com"),
		pcli.CreateIntFlag("bosh-port", "the port of your bosh director", "25555"),
		pcli.CreateStringFlag("bosh-user", "the username for your bosh director", "bosh"),
		pcli.CreateStringFlag("bosh-pass", "the pasword for your bosh director"),
		pcli.CreateBoolFlag("ssl-ignore", "ingore ssl self signed cert warnings"),
		pcli.CreateBoolFlag("print-manifest", "if you would simply like to output a manifest the set this flag as true."),
	}
}

// CloudConfigAction is the action that is executed for
// each cloud config command
func CloudConfigAction(c *cli.Context, cc cloudconfig.CloudConfigDeployer) error {
	manifest := cc.GetCloudConfig(c.Args().Slice())
	lo.G.Debug("we found a manifest and context: ", manifest, c)
	if c.Bool("print-manifest") {
		UIPrint(string(manifest))
		return nil
	}
	bc := getBoshClient(c)
	return bc.PushCloudConfig(manifest)
}

// ProductAction is the action that is executed for each product command
func ProductAction(c *cli.Context, productDeployment product.ProductDeployer) error {
	bc := getBoshClient(c)
	ccm, err := bc.GetCloudConfig()

	if err != nil {
		return err
	}
	bytes, err := ccm.Bytes()

	if err != nil {
		return err
	}

	var manifest []byte
	if manifest, err = productDeployment.GetProduct(c.Args().Slice(), bytes); err != nil {
		lo.G.Errorf("there was an error calling get product: '%v' - '%v'", err.Error(), err)
		return err
	}

	if manifest, err = decorateDeploymentWithBoshUUID(manifest, bc); err == nil {

		if c.Bool("print-manifest") {
			UIPrint(string(manifest))
			return nil

		} else {
			var task enamlbosh.BoshTask
			task, err = uploadProductDeployment(bc, manifest, true)
			lo.G.Debug("bosh task: ", task)
		}
	}

	if err != nil {
		lo.G.Error("there was an error: ", err.Error())
	}
	return err
}

func uploadProductDeployment(client *enamlbosh.Client, manifest []byte, poll bool) (enamlbosh.BoshTask, error) {
	dm := enaml.NewDeploymentManifest(manifest)
	uploadRemoteBoshAssets(dm, client, poll)

	UIPrint("Uploading product deployment...")
	task, err := client.PostDeployment(*dm)
	if err != nil {
		lo.G.Error(err.Error())
		return task, err
	}
	UIPrint("upload complete")
	lo.G.Debug("res: ", task, err)
	err = checkTaskStatus(task, client, poll)
	return task, err
}

// uploadRemoteBoshAssets uploads both stemcells and releases
func uploadRemoteBoshAssets(dm *enaml.DeploymentManifest, boshClient *enamlbosh.Client, poll bool) (err error) {
	var errStemcells error
	var errReleases error
	var remoteStemcells []enaml.Stemcell
	defer UIPrint("remote asset check complete.")
	UIPrint("Checking product deployment for remote assets...")

	if remoteStemcells, err = stemcellsToUpload(dm.Stemcells, boshClient); err == nil {
		if errStemcells = uploadRemoteStemcells(remoteStemcells, boshClient, poll); errStemcells != nil {
			lo.G.Info("issues processing stemcell: ", errStemcells)
		}
	}

	if errReleases = uploadRemoteReleases(dm.Releases, boshClient, poll); errReleases != nil {
		lo.G.Info("issues processing release: ", errReleases)
	}

	if errReleases != nil || errStemcells != nil {
		err = fmt.Errorf("stemcell err: %v   release err: %v", errStemcells, errReleases)
	}
	return
}

func uploadRemoteStemcells(stemcells []enaml.Stemcell, client *enamlbosh.Client, poll bool) error {
	UIPrint("Checking for remote stemcells...")
	defer UIPrint("remote stemcells complete")

	for _, stemcell := range stemcells {
		isRemote := stemcell.URL != "" && stemcell.SHA1 != ""
		if !isRemote {
			UIPrint(fmt.Sprintf("Stemcell %s [%s] already exists", stemcell.Name, stemcell.Version))
			continue
		}
		task, err := client.PostRemoteStemcell(stemcell)
		if err != nil {
			lo.G.Errorf("error uploading stemcell %s [%s]", stemcell.Name, stemcell.Version)
			return err
		}
		lo.G.Debug("task: ", task)
		err = checkTaskStatus(task, client, poll)
		if err != nil {
			return err
		}
	}
	return nil
}

func uploadRemoteReleases(releases []enaml.Release, client *enamlbosh.Client, poll bool) error {
	UIPrint("Checking for remote releases")
	defer UIPrint("remote releases complete")

	for _, release := range releases {
		isRemote := release.URL != "" && release.SHA1 != ""
		if !isRemote {
			continue
		}
		task, err := client.PostRemoteRelease(release)
		if err != nil {
			return err
		}
		lo.G.Debug("task: ", task)
		err = checkTaskStatus(task, client, poll)
		if err != nil {
			return err
		}
	}
	return nil
}

func stemcellsToUpload(stemcells []enaml.Stemcell, client *enamlbosh.Client) ([]enaml.Stemcell, error) {
	result := make([]enaml.Stemcell, 0, len(stemcells))
	for _, stemcell := range stemcells {
		isRemote := stemcell.URL != "" && stemcell.SHA1 != ""
		if !isRemote {
			continue
		}
		exists, err := client.CheckRemoteStemcell(stemcell)
		if err != nil {
			return nil, err
		}
		if !exists {
			result = append(result, stemcell)
		}
	}
	return result, nil
}

func decorateDeploymentWithBoshUUID(deployment []byte, client *enamlbosh.Client) ([]byte, error) {
	var boshinfo *enamlbosh.BoshInfo
	var dm *enaml.DeploymentManifest
	var err error

	if boshinfo, err = client.GetInfo(); err == nil {
		dm = enaml.NewDeploymentManifest(deployment)
		lo.G.Debug("setting uuid on deployment from bosh: ", boshinfo.UUID)
		dm.SetDirectorUUID(boshinfo.UUID)
	}
	return dm.Bytes(), err
}

func checkTaskStatus(task enamlbosh.BoshTask, client *enamlbosh.Client, poll bool) error {
	switch task.State {
	case enamlbosh.StatusCancelled, enamlbosh.StatusError:
		return fmt.Errorf("task is in failed state: %#v", task)
	default:
		if poll {
			return pollTaskAndWait(task, client, -1)
		}
		switch task.State {
		case enamlbosh.StatusCancelled, enamlbosh.StatusError:
			return fmt.Errorf("%s - %s", task.State, task.Description)
		default:
			return nil
		}
	}
}

// pollTaskAndWait will poll the given task until its status is cancelled, done, or error.
// A -1 value for tries indicates infinite attempts.
func pollTaskAndWait(task enamlbosh.BoshTask, client *enamlbosh.Client, tries int) error {
	UIPrint("polling task...")
	defer UIPrint(fmt.Sprintf("Finished with Task %s", task.Description))
	ticker := time.Tick(time.Second)
	count := 0
	for {
		<-ticker
		var err error
		task, err = client.GetTask(task.ID)
		if err != nil {
			return err
		}
		switch task.State {
		case enamlbosh.StatusDone:
			UIPrintStatus(fmt.Sprintf("task state %s", task.State))
			return nil
		case enamlbosh.StatusCancelled, enamlbosh.StatusError:
			err := fmt.Errorf("%s - %s", task.State, task.Description)
			lo.G.Error("task error: " + err.Error())
			return err
		default:
			UIPrintStatus(fmt.Sprintf("task '%s' is %s", task.Description, task.State))
		}
		count++

		if tries != -1 && count >= tries {
			UIPrintStatus("hit poll limit, exiting task poller without error")
			return nil
		}
	}
}

var processingState = 0
var processingStateValues = []string{
	"|", "/", "-", "\\",
}

func getProcessing() string {
	processingState++
	return " [ " + processingStateValues[(processingState%4)] + " ] "
}
