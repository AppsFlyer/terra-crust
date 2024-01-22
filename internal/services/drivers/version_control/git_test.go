package version_control_test

import (
	"fmt"
	"os"
	"testing"

	logger "github.com/AppsFlyer/go-logger"
	version_control "github.com/AppsFlyer/terra-crust/internal/services/drivers/version_control"
)

const (
	TerraCrustModuleName = "terracrust"
	NamingModuleName     = "terraform-aws-resource-naming"
	ZonesModuleName      = "zones"
	TerraCrustURL        = "https://github.com/AppsFlyer/terra-crust.git"
)

var ModulesTestPath = "./temp-git-test"

var mockBadUrl = map[string]*version_control.RemoteModule{
	TerraCrustModuleName: {
		Url: "https://github.com/appsflyer/terra-crust/test/bad",
	},
}

var mockBadVersion = map[string]*version_control.RemoteModule{
	NamingModuleName: {
		Url:     "https://github.com/fajrinazis/terraform-aws-resource-naming.git",
		Version: "bad-tag",
	},
}

func TestCloneAndCleanupInternalGit(t *testing.T) {
	CloneAndCleanup(t, false)
}

func TestCloneAndCleanupExternalGit(t *testing.T) {
	CloneAndCleanup(t, true)
}

func CloneAndCleanupModules(modules map[string]*version_control.RemoteModule, externalGit bool) error {
	log := logger.NewSimple()

	gitDriver := version_control.InitGitProvider(log)

	err := gitDriver.CloneModules(modules, ModulesTestPath, externalGit)
	cErr := gitDriver.CleanModulesFolders(modules, ModulesTestPath)
	if err != nil {
		if cErr != nil {
			return fmt.Errorf("failed to clone and cleanup module %v %v", err, cErr)
		}
		return err
	}
	return nil
}
func CloneAndCleanup(t *testing.T, externalGit bool) {
	var mockModules = map[string]*version_control.RemoteModule{
		TerraCrustModuleName: {
			Url: TerraCrustURL,
		},
		NamingModuleName: {
			Url:     "https://github.com/fajrinazis/terraform-aws-resource-naming.git",
			Version: "v0.3.0",
		},
		ZonesModuleName: {
			Url:  "https://github.com/terraform-aws-modules/terraform-aws-route53.git",
			Path: "modules/zones",
		},
	}
	log := logger.NewSimple()

	gitDriver := version_control.InitGitProvider(log)

	err := gitDriver.CloneModules(mockModules, ModulesTestPath, externalGit)
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
func TestFailBadUrlInternalGit(t *testing.T) {
	FailBadUrl(t, false)
}

func TestFailBadUrlExternalGit(t *testing.T) {
	FailBadUrl(t, true)
}

func TestFailBadVersionInternalGit(t *testing.T) {
	FailBadVersion(t, false)
}

func TestFailBadVersionExternalGit(t *testing.T) {
	FailBadVersion(t, true)
}

func FailBadUrl(t *testing.T, externalGit bool) {
	t.Parallel()

	err := CloneAndCleanupModules(mockBadUrl, externalGit)
	if err == nil {
		t.Errorf("expected error received error nil")
	}
}

func FailBadVersion(t *testing.T, externalGit bool) {
	t.Parallel()

	err := CloneAndCleanupModules(mockBadVersion, externalGit)
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
