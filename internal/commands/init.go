package commands

import (
	"fmt"

	"gpm/internal/project"
)

func InitProject(args []string) {
	if len(args) < 1 {
		fmt.Println("Please specify project name")
		return
	}

	projectName := args[0]
	err := project.Create(projectName)
	if err != nil {
		fmt.Printf("Error creating project: %v\n", err)
		return
	}

	fmt.Printf("Project %s created successfully!\n", projectName)
}
