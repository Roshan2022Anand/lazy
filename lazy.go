package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/google/go-github/v50/github"
	"golang.org/x/oauth2"
)

const projectsPath = "C:/hackathon/program"

// type Template struct {
// 	cmd          string
// 	folderNeeded bool
// }

// type Templates map[string]Template

// var templates = Templates{
// 	"--cpp": {
// 		cmd:          "mkdir bin && mkdir src && mkdir user && echo //Code Here > main.cpp",
// 		folderNeeded: false,
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
	// Handling Invalid arguments
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

func directoryExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func createGitHubRepo(token, name string) (*github.Repository, error) {
	// Create authenticated GitHub client
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// Create repository request
	repo := &github.Repository{
		Name:    github.String(name),
		Private: github.Bool(false),
	}

	// Create the repository
	createdRepo, _, err := client.Repositories.Create(ctx, "", repo)
	if err != nil {
		return nil, err
	}
	return createdRepo, nil
}

func connectLocalToRemote(localPath, remoteURL string) error {
	cmd := exec.Command("git", "remote", "add", "origin", remoteURL)
	cmd.Dir = localPath
	return cmd.Run()
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
