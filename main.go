package main

import (
	"fmt"
	"os"

	gen "github.com/munjeli/ezapi/generators"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("missing positional arguments: `ezapi api apiName targetDirectory`")
		os.Exit(1)
	}
	genType := os.Args[1]
	name := os.Args[2]
	targetDirectory := os.Args[3]
	switch genType {
	case "api":
		gen.GenerateAPI(name, targetDirectory)
	default:
		fmt.Printf("%v is not a supported generator.\nSupported generator is `api`.", genType)
	}
}
