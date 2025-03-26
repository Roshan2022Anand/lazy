package main

import (
	"fmt"
	"os"
)

// type Template struct {
// 	cmd          string
// 	folderNeeded bool
// }

// type Templates map[string]Template

// var templates = Templates{
// 	"--cpp": {
// 		cmd:          "mkdir bin && mkdir src && mkdir user && echo //Code Here > main.cpp",
// 		folderNeeded: false, // Changed from = to : for struct literal
// 	},
// }

func displayHelp() {
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("	create	- create a new project")
	fmt.Println("	list	- lists local projects")
	fmt.Println("Templates:")
	fmt.Println("	--cpp	- basic template for a C++ project")
	fmt.Println()
}

// func listProjects() {

// }

func main() {
	argumentsCount := len(os.Args) - 1
	if argumentsCount > 1 {
		if os.Args[1] == "create" {
			// createNewProject()
		} else if os.Args[1] == "list" {
			// listProjects()
		} else {
			displayHelp()
		}
	} else {
		displayHelp()
	}
}
