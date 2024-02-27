package reporender

import (
	"fmt"
	"os"
)

// Renders the gitops Repo https://github.com/PatrickLaabs/argo-hub-template
//
// This Repo contains some placeholder strings.
// To be more precise:
// - GITHUB_USER
// - GITHUB_USER_EMAIL

// FullStage combines everything, that is needed, to fully prepare the gitops repo for the end-user
func FullStage() {
	fmt.Println("Rendering the gitops template repo")
}

// retrieveEnvs retrieves and reads the os.Env variables needed for further preperation
// GITHUB_USER
// GITHUB_USER_EMAIL
// These Envs are mandatory to run this CLI and will be checked on cluster-creation runtime.
// At this point, they will be available
func retrieveEnvs() (string, string, error) {
	// Get GITHUB_USERNAME environment var
	var username string
	var usermail string

	if os.Getenv("GITHUB_USERNAME") == "" {
		fmt.Println("Missing Github Username, please set it. Exiting now.")
		os.Exit(1)
	} else {
		username = os.Getenv("GITHUB_USERNAME")
	}

	if os.Getenv("GITHUB_USER_EMAIL") == "" {
		fmt.Println("Missing Github User Email, please set it. Exiting now.")
		os.Exit(1)
	} else {
		usermail = os.Getenv("GITHUB_USER_EMAIL")
	}

	return username, usermail, nil
}

// gitClone clones the gitops template repo from github, to the local working directory
func gitClone() {

}

// gitConfigs configures the local git repo (user, mail, etc)
func gitConfigs() {}

// replaceStrings replaces the Placeholder strings inside the gitops repo
func replaceStrings() {
	username, usermail, err := retrieveEnvs()
	if err != nil {
		fmt.Println("Error retrieving token:", err)
		os.Exit(1)
	}

	fmt.Println(username)
	fmt.Println(usermail)

	// we need to look over many files, and have to replace two different kind of Placeholder strings

}

// gitPush pushes the changes to the users github repository
func gitPush() {}
