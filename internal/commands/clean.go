package commands

import (
	"fmt"
	"os"
	"path/filepath"
)

func CleanProject(args []string) {
	// Удаляем бинарные файлы
	if err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && (info.Name() == "goproj" || info.Name() == "main") {
			return os.Remove(path)
		}
		return nil
	}); err != nil {
		fmt.Printf("Clean failed: %v\n", err)
		return
	}

	fmt.Println("Project cleaned successfully!")
}
