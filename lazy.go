package main

import (
	"fmt"
	"lazy/nodesetup"
	"os"
	"sync"
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

	
	wg := sync.WaitGroup{}
	wg.Add(3)
	var name string
	fmt.Print("Enter your Name : ")
	fmt.Scanln(&name)
	err := os.Mkdir(name, 0755)
	if err != nil {
		fmt.Println("Error creating directory", err)
	}
	os.Chdir(name)

	go nodesetup.CtreateNodeServer(&wg)
	go nodesetup.InitNode(&wg)
	go nodesetup.CreateFolderStruct(&wg)

	wg.Wait()
	fmt.Println("Project setup completed")
}
