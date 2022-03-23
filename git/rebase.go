package git

import (
	"os"
	"os/exec"
)

// Rebase executes 'git rebase -i <commit_hash>' with the supplied commit hash
func Rebase(hash string) error {
	cmd := exec.Command("git", "rebase", "-i", hash)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
