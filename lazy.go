package main

import (
	"fmt"
	"os"
)

const projectsPath = "C:/test/programs"

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

func createNewProject() {
	fmt.Println("Creating")
}

func listProjects(argumentsCount int) {
	if argumentsCount > 1 || argumentsCount < 1 {
		fmt.Println("\nInvalid Arguments!!")
		fmt.Println("Usage: lazy list")
	} else {
		entries, _ := os.ReadDir(projectsPath)
		fmt.Println("\nDirectories:")
		for _, e := range entries {
			if e.IsDir() {
				fmt.Println(" ", e.Name())
			}
		}
	}
}

func main() {
	argumentsCount := len(os.Args) - 1
	if argumentsCount > 0 {
		if os.Args[1] == "create" {
			createNewProject()
		} else if os.Args[1] == "list" {
			listProjects(argumentsCount)
		} else {
			displayHelp()
		}
	} else {
		displayHelp()
	}
}
