package generators

import (
	"fmt"
	"os"
)

// GenerateAPI will make the directories then use
// templates to stub out CRUD and a server.
func GenerateAPI(name, targetDir string) error {
	err := makeDirs(name, targetDir, "api")
	if err != nil {
		return err
	}
	return nil
}

// makeAPIDirs makes the directories for the API type.
func makeDirs(name, targetDir, apiType string) error {
	var dirs []string
	if apiType == "api" {
		dirs = []string{fmt.Sprintf("%s/apis/%s/server", targetDir, name)}
	} else if apiType == "netsrv" {
		dirs = []string{
			fmt.Sprintf("%s/service/%s", targetDir, name),
			fmt.Sprintf("%s/server", targetDir),
		}
	} else {
		return fmt.Errorf("no correct API type for %v", name)
	}
	for _, path := range dirs {
		err := os.MkdirAll(path, 0775)
		if err != nil {
			return err
		}
		fmt.Printf("Created directories: %s\n", path)
	}
	return nil
}

// GenerateNetworkedService generates stubs for a service with a
// single endpoint and server for exposing a utility on a server.
// This is not what to use for a user-facing application, instead it is
// for internal services that don't need full CRUD exposed, like daemons.
func GenerateNetworkedService(name, targetDir string) error {
	err := makeDirs(name, targetDir, "netsrv")
	if err != nil {
		return err
	}
	return nil
}
