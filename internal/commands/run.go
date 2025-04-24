package commands

import (
	"fmt"
	"os"
	"os/exec"
)

func RunProject(args []string) {
	cmd := exec.Command("go", "run", "./cmd")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Run failed: %v\n", err)
	}
}
