package version_control_test

import (
	logger "github.com/AppsFlyer/go-logger"
	version_control "github.com/AppsFlyer/terra-crust/internal/services/drivers/version-control"
	"os"
	"testing"
)

const (
	TerraCrustModuleName = "terracrust"
	TerraCrustURL        = "https://github.com/AppsFlyer/terra-crust"
	ModulesTestPath      = "./temp-git-test"
)

var mockModules = map[string]*version_control.RemoteModule{
	TerraCrustModuleName: {
		Url: TerraCrustURL,
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

	if len(folders) != 1 {
		t.Errorf("Expected 1 folder count received %d", len(folders))
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

	if len(folders) != 0 {
		t.Errorf("Expected 0 folder count received %d", len(folders))
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
