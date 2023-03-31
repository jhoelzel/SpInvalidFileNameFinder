package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	renameEnabled bool
	dir           string
)

func main() {
	flag.BoolVar(&renameEnabled, "rename", false, "Enable renaming of files and folders with invalid names")
	flag.StringVar(&dir, "dir", "", "Specify the directory to process (defaults to the current working directory)")
	flag.Parse()

	var err error

	if dir == "" {
		dir, err = os.Getwd()
		if err != nil {
			log.Fatalf("Error getting current directory: %v", err)
		}
	}

	logFile := filepath.Join(dir, "invalid_filenames.log")

	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Error creating log file: %v", err)
	}
	defer file.Close()

	log.SetOutput(file)

	err = filepath.Walk(dir, processPath)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

// processPath checks the name of the file or folder and renames it if it contains invalid characters.
func processPath(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	// Skip files and folders that start with a tilde
	if strings.HasPrefix(info.Name(), "~") {
		if info.IsDir() {
			return filepath.SkipDir
		}
		return nil
	}
	// Skip files and folders that start with a dot
	if strings.HasPrefix(info.Name(), ".") {
		if info.IsDir() {
			return filepath.SkipDir
		}
		return nil
	}

	if isNameInvalid(info.Name()) {
		log.Println("Invalid name:", path)
		newName := replaceInvalidCharacters(info.Name())
		newPath := filepath.Join(filepath.Dir(path), newName)
		log.Println("Suggested name:", newPath)

		if renameEnabled {
			err := os.Rename(path, newPath)
			if err != nil {
				log.Printf("Error renaming file: %v", err)
			}
		}
	}
	return nil
}

// isNameInvalid checks if the given name contains any invalid characters for SharePoint.
func isNameInvalid(name string) bool {
	invalidCharacters := regexp.MustCompile(`[\\/:*?"<>|#]|^\.`)
	return invalidCharacters.MatchString(name)
}

// replaceInvalidCharacters replaces invalid characters in the given name with suitable replacements.
func replaceInvalidCharacters(name string) string {
	name = strings.ReplaceAll(name, "#", "NR")
	invalidCharacters := regexp.MustCompile(`[\\/:*?"<>|]`)
	return invalidCharacters.ReplaceAllString(name, "_")
}
