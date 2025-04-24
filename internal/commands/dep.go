package commands

import (
	"fmt"
	"os"
	"os/exec"
)

func ManageDependencies(args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: goproj dep [add|remove] <package>")
		return
	}

	action := args[0]
	pkg := args[1]

	var cmd *exec.Cmd

	switch action {
	case "add":
		cmd = exec.Command("go", "get", pkg)
	case "remove":
		cmd = exec.Command("go", "mod", "edit", "-droprequire="+pkg)
	default:
		fmt.Printf("Unknown dependency action: %s\n", action)
		return
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Dependency operation failed: %v\n", err)
		return
	}

	fmt.Printf("Dependency %s %sed successfully\n", pkg, action)
}
