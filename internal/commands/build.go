package commands

import (
	"fmt"
	"os"
	"os/exec"
)

func BuildProject(args []string) {
	cmd := exec.Command("go", "build", "./...")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Build failed: %v\n", err)
		return
	}

	fmt.Println("Build successful!")
}
