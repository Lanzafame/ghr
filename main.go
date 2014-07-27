package main

import (
	"fmt"
	"log"
	"os"
)

type Info struct {
	Token           string
	TagName         string
	RepoName        string
	OwnerName       string
	TargetCommitish string
	Draft           bool
	Prerelease      bool
}

func debug(v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		log.Println(v...)
	}
}

func NewInfo() Info {
	return Info{
		TargetCommitish: "master",
		Draft:           false,
		Prerelease:      false,
	}
}

func main() {
	// call ghrMain in a separate function
	// so that it can use defer and have them
	// run before the exit.
	os.Exit(ghrMain())
}

func ghrMain() int {

	if os.Getenv("GITHUB_TOKEN") == "" {
		fmt.Fprintf(os.Stderr, "Please set your Github API Token in the GITHUB_TOKEN env var\n")
		return 1
	}

	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: ghr <tag> <artifact>\n")
		return 1
	}

	tag := os.Args[1]
	debug(tag)

	artifacts := os.Args[2]
	debug(artifacts)

	// git config --global user.name
	// tcnksm
	owner, err := GitOwner()
	if err != nil || owner == "" {
		fmt.Fprintf(os.Stderr, "Please set `git config --global user.name`\n")
		return 1
	}
	debug(owner)

	// git config --local remote.origin.url
	// https://github.com/tcnksm/ghr.git
	remoteURL, err := GitRemote()
	if err != nil || remoteURL == "" {
		fmt.Fprintf(os.Stderr, "Please set remote host of your project\n")
		return 1
	}
	debug(remoteURL)

	repo := GitRepoName(remoteURL)
	debug(repo)

	info := NewInfo()
	info.Token = os.Getenv("GITHUB_TOKEN")
	info.TagName = tag
	info.OwnerName = owner
	info.RepoName = repo
	debug(info)

	id, err := GetReleaseID(info)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return 1
	}
	debug(id)

	if id == -1 {
		err = CreateNewRelease(info)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			return 1
		}

		id, err = GetReleaseID(info)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			return 1
		}
	}

	debug(id)
	//

	return 0
}
