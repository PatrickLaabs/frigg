package reporender

import (
	"fmt"
	"github.com/PatrickLaabs/frigg/pkg/common/wait"
	"github.com/fatih/color"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"time"
)

// Renders the gitops Repo https://github.com/PatrickLaabs/argo-hub-template
//
// This Repo contains some placeholder strings.
// To be more precise:
// - GITHUB_USER
// - GITHUB_MAIL

// FullStage combines everything, that is needed, to fully prepare the gitops repo for the end-user
func FullStage() {
	println(color.GreenString("Rendering the gitops template repo"))

	username, err := retrieveGithubUserEnv()
	if err != nil {
		fmt.Printf("Error retrieving username: %v", err)
	}

	usermail, err := retrieveGithubUserMailEnv()
	if err != nil {
		fmt.Printf("Error retrieving mail: %v", err)
	}
	homedir, _ := os.UserHomeDir()
	friggDirName := ".frigg"
	friggDir := homedir + "/" + friggDirName

	repoName := "argo-hub-test"
	localRepoStoragePath := friggDir + "/" + repoName

	githubLogin()
	gitCreateFromTemplate()
	wait.Wait(5 * time.Second)
	gitClone()
	err = replaceStrings(localRepoStoragePath, username, usermail)
	if err != nil {
		return
	}
	gitCommit()
	gitPush()
}

// retrieveGithubUserEnv retrieves and reads the os.Env variables needed for further preperation
// GITHUB_USER
func retrieveGithubUserEnv() (string, error) {
	// Get GITHUB_USERNAME environment var
	var username string

	if os.Getenv("GITHUB_USERNAME") == "" {
		println(color.RedString("Missing Github Username, please set it. Exiting now."))
		os.Exit(1)
	} else {
		username = os.Getenv("GITHUB_USERNAME")
	}

	return username, nil
}

// retrieveGithubUserMailEnv retrieves and reads the os.Env variables needed for further preperation
// GITHUB_MAIL
func retrieveGithubUserMailEnv() (string, error) {
	var usermail string

	if os.Getenv("GITHUB_MAIL") == "" {
		println(color.RedString("Missing Github User Email, please set it. Exiting now."))
		os.Exit(1)
	} else {
		usermail = os.Getenv("GITHUB_MAIL")
	}

	return usermail, nil
}

// githubLogin logs in to github via github cli using the provided github token
func githubLogin() {
	println(color.GreenString("Loggin in to Github with your provided Github Token"))

	cmd := exec.Command("gh", "auth", "login")

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		println(color.RedString("Error: %v", err))
		println(color.YellowString(string(output)))
		return
	}
	fmt.Println(string(output))
}

// gitCreateFromTemplate creates a repository based from the template repo 'argo-hub-template'
func gitCreateFromTemplate() {
	fmt.Println("Creating Argohub Repo out of Template Repo")

	username, err := retrieveGithubUserEnv()
	if err != nil {
		fmt.Printf("Error retrieving token: %v", err)
	}
	fmt.Println(username)

	repoName := "argo-hub-test"
	targetRepoName := username + "/" + repoName

	cmd := exec.Command("gh", "repo", "create",
		targetRepoName, "--private",
		"--template=PatrickLaabs/argo-hub-template",
	)

	// Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running clusterctl: %s\n", err)
		fmt.Println(string(output))
		return
	}
	fmt.Println(string(output))
}

// gitClone clones the gitops template repo from github, to the local working directory
func gitClone() {
	fmt.Println("Cloning the new repository to the local working directory")

	username, err := retrieveGithubUserEnv()
	if err != nil {
		fmt.Printf("Error retrieving token: %v", err)
	}

	homedir, _ := os.UserHomeDir()
	friggDirName := ".frigg"
	friggDir := homedir + "/" + friggDirName

	repoName := "argo-hub-test"
	// git@github.com:PatrickLaabs/frigg.git
	repoUrl := "git@github.com:" + username + "/" + repoName + ".git"
	localRepoStoragePath := friggDir + "/" + repoName

	_, err = git.PlainClone(localRepoStoragePath, false, &git.CloneOptions{
		URL:      repoUrl,
		Progress: os.Stdout,
	})
	if err != nil {
		fmt.Printf("Error cloning your Argohub Repo: %v\n", err)
	}
}

// replaceStrings replaces the Placeholder strings inside all files in the gitops repo
func replaceStrings(dirPath string, username string, usermail string) error {
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		reGhUser := regexp.MustCompile(`GITHUB_USERNAME`)
		reGhMail := regexp.MustCompile(`GITHUB_MAIL`)

		// Replace GITHUB_USER and GITHUB_MAIL
		newdata := replaceInString(data, reGhUser, username)
		newdata = replaceInString(newdata, reGhMail, usermail)

		// Open the file for writing and replace content
		file, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0644) // Adjust permissions as needed
		if err != nil {
			return err
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {

			}
		}(file)

		_, err = file.Write(newdata)
		return err
	})

	return err
}

// replaceInString replaces specific pattern with a new string
func replaceInString(data []byte, re *regexp.Regexp, replacement string) []byte {
	return re.ReplaceAll(data, []byte(replacement))
}

// gitCommit commits local changes
func gitCommit() {
	fmt.Println("Committing local changes")

	username, err := retrieveGithubUserEnv()
	if err != nil {
		fmt.Printf("Error retrieving token: %v", err)
	}
	fmt.Println(username)

	usermail, err := retrieveGithubUserMailEnv()
	if err != nil {
		fmt.Printf("Error retrieving token: %v", err)
	}
	fmt.Println(usermail)

	homedir, _ := os.UserHomeDir()
	friggDirName := ".frigg"
	friggDir := homedir + "/" + friggDirName

	repoName := "argo-hub-test"
	localRepoStoragePath := friggDir + "/" + repoName

	// Opens an already existing repository.
	r, err := git.PlainOpen(localRepoStoragePath)
	if err != nil {
		fmt.Printf("error on accessing the local repo: %v\n", err)
	}

	w, err := r.Worktree()
	if err != nil {
		return
	}

	_, err = w.Add(".")
	if err != nil {
		fmt.Printf("Error on committing local changes: %v\n", err)
	}

	// We can verify the current status of the worktree using the method Status.
	status, err := w.Status()
	if err != nil {
		fmt.Printf("Error on checking the status: %v\n", err)
	}
	fmt.Println(status)

	// Commits the current staging area to the repository, with the new file
	// just created. We should provide the object.Signature of Author of the
	// commit Since version 5.0.1, we can omit the Author signature, being read
	// from the git config files.
	commit, err := w.Commit("example go-git commit", &git.CommitOptions{
		Author: &object.Signature{
			Name:  username,
			Email: usermail,
			When:  time.Now(),
		},
	})
	if err != nil {
		fmt.Printf("Error on committing: %v\n\n", err)
	}

	// Prints the current HEAD to verify that all worked well.
	obj, err := r.CommitObject(commit)
	if err != nil {
		fmt.Printf("Error checking the status: %v\n", err)
	}

	fmt.Println(obj)
}

// gitPush pushes the changes to the users github repository
func gitPush() {
	fmt.Println("Pushing local changes to the remote repo")

	homedir, _ := os.UserHomeDir()
	friggDirName := ".frigg"
	friggDir := homedir + "/" + friggDirName

	repoName := "argo-hub-test"
	localRepoStoragePath := friggDir + "/" + repoName

	// Opens an already existing repository.
	r, err := git.PlainOpen(localRepoStoragePath)
	if err != nil {
		fmt.Printf("error on accessing the local repo: %v\n", err)
	}

	err = r.Push(&git.PushOptions{})
	if err != nil {
		fmt.Printf("Error on pushing: %v\n", err)
	}
}
