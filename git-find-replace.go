package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"sort"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Usage: %s <search string> <replacement string>\n", os.Args[0])
	}

	searchString := os.Args[1]
	replacementString := os.Args[2]

	// Running "git grep -l -F" to find a list of files containing instances of searchString
	cmd := exec.Command("git", "grep", "-l", "-F", searchString)
	cmdOutput, err := cmd.Output()
	if err != nil {
		log.Fatalf("Failed to execute git grep: %s\n", err)
	}

	// Splitting command output into a slice of file paths
	filePaths := strings.Split(string(cmdOutput), "\n")
	sort.Strings(filePaths)
	if len(filePaths) != 0 && filePaths[0] == "" {
		filePaths = filePaths[1:]
	}

	fmt.Fprintf(os.Stderr, "Replacing in %d files...\n", len(filePaths))
	for i, filePath := range filePaths {
		if filePath == "" {
			continue
		}

		// Reading file content
		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			log.Fatalf("Failed to read file %s: %s\n", filePath, err)
		}

		// Replacing instances of searchString with replacementString
		newContent := strings.ReplaceAll(string(content), searchString, replacementString)

		// Writing new content back to file
		err = ioutil.WriteFile(filePath, []byte(newContent), 0)
		if err != nil {
			log.Fatalf("Failed to write to file %s: %s\n", filePath, err)
		}

		fmt.Fprintf(os.Stderr, "[%d/%d]: replaced in %s\n", i+1, len(filePaths), filePath)
	}
}
