package git

import (
	"bytes"
	"os/exec"
)

// A Structure which contains all the info of a commit
type Commit struct {
	Hash        string // The hash of the Commit
	Date        string // The author date of the Commit
	Description string // The Description of the Commit
	Author      string // The author of this Commit
}

// CommitHistory Fetches the commit history known the the working branch and
// returns all found commits in a slice of Commit structs.
func CommitHistory() ([]Commit, error) {
	output, err := exec.Command("git", "log", "--pretty=format:%h,%ad,%s,%an", "--date=short", "-50").Output()

	if err != nil {
		return nil, err
	}

	var commits []Commit

	commitBytes := bytes.Split(output, []byte("\n"))

	for _, commit := range commitBytes {
		if len(commit) == 0 {
			continue
		}

		commitInfo := bytes.Split(commit, []byte(","))

		commits = append(commits, Commit{
			Hash:        string(commitInfo[0]),
			Date:        string(commitInfo[1]),
			Description: string(commitInfo[2]),
			Author:      string(commitInfo[3]),
		})
	}

	return commits, nil
}
