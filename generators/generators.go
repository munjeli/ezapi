package generators

import (
	"fmt"
	"os"
	"text/template"

	tmpl "github.com/munjeli/ezapi/templates"
)

type templateInput struct {
	Name      string
	TargetDir string
}

// GenerateAPI will make the directories then use
// templates to stub out CRUD and a server.
func GenerateAPI(name, targetDir string) error {
	err := makeDirs(name, targetDir, "api")
	if err != nil {
		return err
	}
	err = generateFilesFromTemplates("api", name, targetDir)
	if err != nil {
		return err
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
	err = generateFilesFromTemplates("netsrv", name, targetDir)
	if err != nil {
		return err
	}
	return nil
}

// makeAPIDirs makes the directories for the API type.
func makeDirs(name, targetDir, apiType string) error {
	var dirs []string
	if apiType == "api" {
		dirs = []string{fmt.Sprintf("%s/apis/%sapi/server", targetDir, name)}
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

// generateFilesFromTemplates will use the API type to select the correct templates
// and generate the files from an input object.
func generateFilesFromTemplates(apiType, name, targetDir string) error {
	i := templateInput{
		Name:      name,
		TargetDir: targetDir,
	}
	if apiType == "api" {
		if err := generateMakefile(i, tmpl.APIMakeTemplate); err != nil {
			return err
		}
		if err := generateAPIFiles(i); err != nil {
			return err
		}
	} else if apiType == "netsrv" {
		if err := generateMakefile(i, tmpl.NetSrvMakeTemplate); err != nil {
			return err
		}
		if err := generateNetSrvFiles(i); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("not a supported API Type: %s", apiType)
	}
	return nil
}

func generateMakefile(i templateInput, tmp string) error {
	t := template.Must(template.New("Makefile").Parse(tmp))
	f, err := os.Create(fmt.Sprintf(tmpl.MakePath, i.TargetDir))
	if err != nil {
		return err
	}
	err = t.Execute(f, i)
	if err != nil {
		return err
	}
	return nil
}

func generateNetSrvFiles(i templateInput) error {
	//for _, tf := range tfiles {
	//	t := template.Must(template.New(tf.Title).Parse(tf.Tmpl))
	//	f, err := os.Create()
	//	if err != nil {
	//		return err
	//	}
	//	err = t.Execute(f, i)
	//	if err != nil {
	//		return err
	//	}
	//}
	return nil
}

func generateAPIFiles(i templateInput) error {
	return nil
}
