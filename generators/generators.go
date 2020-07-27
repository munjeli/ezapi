package generators

import (
	"fmt"
	"os"
)

// GenerateAPI will make the directories then use
// templates to stub out CRUD and a server.
func GenerateAPI(name, targetDir string) error {
	err := MakeAPIDirs(name, targetDir)
	if err != nil {
		return err
	}
	return nil
}

// MakeAPIDirs makes the directories for a new API.
func MakeAPIDirs(name, targetDir string) error {
	dirPath := fmt.Sprintf("%s/apis/%s/server", targetDir, name)
	fmt.Printf("Creating directories for new API: %s\n", dirPath)
	err := os.MkdirAll(dirPath, 0775)
	if err != nil {
		return err
	}
	return nil
}
