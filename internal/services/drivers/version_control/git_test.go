package version_control_test

import (
	"os"
	"testing"

	logger "github.com/AppsFlyer/go-logger"
	version_control "github.com/AppsFlyer/terra-crust/internal/services/drivers/version_control"
)

const (
	TerraCrustModuleName = "terracrust"
	NamingModuleName     = "terraform-aws-resource-naming"
	ZonesModuleName      = "zones"
	TerraCrustURL        = "https://github.com/AppsFlyer/terra-crust"

	ModulesTestPath = "./temp-git-test"
)

var mockModules = map[string]*version_control.RemoteModule{
	TerraCrustModuleName: {
		Url: TerraCrustURL,
	},
	NamingModuleName: {
		Url:     "https://github.com/fajrinazis/terraform-aws-resource-naming.git",
		Version: "v0.23.1",
	},
	ZonesModuleName: {
		Url:  "https://github.com/terraform-aws-modules/terraform-aws-route53.git",
		Path: "modules/zones",
	},
}
var mockBadModules = map[string]*version_control.RemoteModule{
	TerraCrustModuleName: {
		Url: "https://github.com/appsflyer/terra-crust/test/bad",
	},
}

func TestCloneAndCleanup(t *testing.T) {
	t.Parallel()
	log := logger.NewSimple()

	gitDriver := version_control.InitGitProvider(log)

	err := gitDriver.CloneModules(mockModules, ModulesTestPath)
	if err != nil {
		t.Errorf(err.Error())
	}

	folders, err := getFolderAsMap(ModulesTestPath)
	if err != nil {
		t.Errorf(err.Error())
	}

	if _, exist := folders[TerraCrustModuleName]; !exist {
		t.Errorf("Clone failed , did not find terra-crust")
	}

	if _, exist := folders[ZonesModuleName]; !exist {
		t.Errorf("Clone failed , did not find zones")
	}

	if _, exist := folders[NamingModuleName]; !exist {
		t.Errorf("Clone failed , did not find naming")
	}

	if len(folders) != 3 {
		t.Errorf("Expected 3 folder count received %d", len(folders))
	}

	if err = gitDriver.CleanModulesFolders(mockModules, ModulesTestPath); err != nil {
		t.Errorf("failed to clean up the downloaded modules,  %s", err.Error())
	}

	folders, err = getFolderAsMap(ModulesTestPath)
	if err != nil {
		t.Errorf(err.Error())
	}

	if _, exist := folders[TerraCrustModuleName]; exist {
		t.Errorf("cleanup-failed terracrust folder still exists")
	}

	if _, exist := folders[ZonesModuleName]; exist {
		t.Errorf("cleanup-failed zones folder still exists")
	}

	if _, exist := folders[NamingModuleName]; exist {
		t.Errorf("cleanup-failed naming folder still exists")
	}

	if len(folders) != 0 {
		t.Errorf("Expected 0 folder count received %d", len(folders))
	}
}

func TestFailBadUrl(t *testing.T) {
	t.Parallel()
	log := logger.NewSimple()
	gitDriver := version_control.InitGitProvider(log)

	err := gitDriver.CloneModules(mockBadModules, ModulesTestPath)
	if err == nil {
		t.Errorf("expected error received error nil")
	}
}

func getFolderAsMap(path string) (map[string]struct{}, error) {
	folders := make(map[string]struct{})

	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			folders[file.Name()] = struct{}{}
		}
	}

	return folders, err
}
