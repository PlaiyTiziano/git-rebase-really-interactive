package main

import (
	"fmt"
	"os"

	"github.com/PlaiyTiziano/git-rebase-really-interactive/git"
	"github.com/ktr0731/go-fuzzyfinder"
)

func getCommitToRebaseOn(commits []git.Commit) (string, error) {
	idx, err := fuzzyfinder.Find(
		commits,
		func(i int) string {
			return fmt.Sprintf("[%s - %s] %s", commits[i].Date, commits[i].Author, commits[i].Description)
		},
	)
	if err != nil {
		return "", err
	}

	return commits[idx].Hash, nil
}

func main() {
	commits, err := git.CommitHistory()
	if err != nil {
		fmt.Println("Failed to execute 'git log --oneline -50'")
		os.Exit(1)
	}

	commitHash, err := getCommitToRebaseOn(commits)
	if err != nil {
		fmt.Printf("Failed to select a commit (%v)\n", err)
		os.Exit(1)
	}

	git.Rebase(commitHash)
}
