package main

import (
	"fmt"
	"os"

	"gpm/internal/commands"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	switch cmd {
	case "init":
		commands.InitProject(args)
	case "build":
		commands.BuildProject(args)
	case "run":
		commands.RunProject(args)
	case "dep":
		commands.ManageDependencies(args)
	case "clean":
		commands.CleanProject(args)
	default:
		fmt.Printf("Unknown command: %s\n", cmd)
		printHelp()
	}
}

func printHelp() {
	fmt.Println("Go Project Manager (goproj)")
	fmt.Println("Usage:")
	fmt.Println("  gpm init <project-name>   - Create a new Go project")
	fmt.Println("  gpm build                - Build the project")
	fmt.Println("  gpm run                  - Run the project")
	fmt.Println("  gpm dep add <pkg>        - Add a dependency")
	fmt.Println("  gpm dep remove <pkg>     - Remove a dependency")
	fmt.Println("  gpm clean                - Clean build artifacts")
}
