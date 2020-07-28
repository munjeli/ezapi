package generators

import (
	"fmt"
	"os"
)

// GenerateAPI will make the directories then use
// templates to stub out CRUD and a server.
func GenerateAPI(name, targetDir string) error {
	err := makeAPIDirs(name, targetDir)
	if err != nil {
		return err
	}
	return nil
}

// makeAPIDirs makes the directories for a new API.
func makeAPIDirs(name, targetDir string) error {
	dirPath := fmt.Sprintf("%s/apis/%s/server", targetDir, name)
	err := os.MkdirAll(dirPath, 0775)
	if err != nil {
		return err
	}
	fmt.Printf("Created directories for new API: %s\n", dirPath)
	return nil
}

// GenerateNetworkedService generates stubs for a service with a
// single endpoint and server for exposing a utility on a server.
// This is not what to use for a user-facing application, instead it is
// for internal services that don't need full CRUD exposed, like daemons.
func GenerateNetworkedService(name, targetDir string) error {
	err := makeNetSrvDirs(name, targetDir)
	if err != nil {
		return err
	}
	return nil
}

// makeNetSrvDirs makes the directories for a networked service.
func makeNetSrvDirs(name, targetDir string) error {
	servicePath := fmt.Sprintf("%s/service/%s", targetDir, name)
	serverPath := fmt.Sprintf("%s/server", targetDir)
	for _, d := range []string{serverPath, servicePath} {
		err := os.MkdirAll(d, 0775)
		if err != nil {
			return err
		}
	}
	fmt.Println("Created directories for new networked service")
	return nil
}
