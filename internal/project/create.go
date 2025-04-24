package project

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
)

type Project struct {
	Name string
	Path string
}

func Create(name string) error {
	if err := os.Mkdir(name, 0755); err != nil {
		return err
	}

	cmd := exec.Command("go", "mod", "init", name)
	cmd.Dir = name
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("go mod init failed: %v", err)
	}

	dirs := []string{
		"cmd",
		"internal",
		"pkg",
		// "api",
		"config",
		"build",
	}

	for _, dir := range dirs {
		if err := os.Mkdir(filepath.Join(name, dir), 0755); err != nil {
			return err
		}
	}

	mainFile := filepath.Join(name, "cmd", "main.go")
	if err := createMainFile(mainFile, name); err != nil {
		return err
	}

	return nil
}

func createMainFile(path, projectName string) error {
	tmpl := `package main

import "fmt"

func main() {
	fmt.Println("Welcome to {{.ProjectName}}!")
}
`

	t, err := template.New("main").Parse(tmpl)
	if err != nil {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	data := struct {
		ProjectName string
	}{
		ProjectName: projectName,
	}

	return t.Execute(f, data)
}
