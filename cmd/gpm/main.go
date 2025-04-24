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
	puts("Go Project Manager (goproj)")
	puts("Usage:")
	puts("  gpm init <project-name>  - Create a new Go project")
	puts("  gpm build                - Build the project")
	puts("  gpm run                  - Run the project")
	puts("  gpm dep add <pkg>        - Add a dependency")
	puts("  gpm dep remove <pkg>     - Remove a dependency")
	puts("  gpm clean                - Clean build artifacts")
}

var puts = func(args ...string) {
	fmt.Println(args)
}
