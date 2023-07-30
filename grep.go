package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func grepFile(pattern string, filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file %s: %s\n", filePath, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(pattern)

	for scanner.Scan() {
		line := scanner.Text()
		if re.MatchString(line) {
			fmt.Printf("%s: %s\n\n", filePath, line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file %s: %s\n", filePath, err)
		return
	}
}

func grepRecursive(pattern string, dirPath string) {
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing %s: %s\n", path, err)
			return nil
		}

		if !info.IsDir() {
			grepFile(pattern, path)
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error walking through directory %s: %s\n", dirPath, err)
	}
}

func grep(pattern string, paths []string, recursive bool) {
	for _, path := range paths {
		fileInfo, err := os.Stat(path)
		if err != nil {
			fmt.Printf("Error accessing %s: %s\n", path, err)
			continue
		}

		if fileInfo.IsDir() {
			if recursive {
				grepRecursive(pattern, path)
			} else {
				fmt.Printf("Error: %s is a directory. Use -R for recursive search.\n", path)
			}
		} else {
			grepFile(pattern, path)
		}
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: grep [-R] <pattern> <file/dir> [<file/dir> ...]")
		os.Exit(1)
	}

	recursive := false
	patternIndex := 1

	if os.Args[1] == "-R" {
		recursive = true
		patternIndex = 2
	}

	pattern := os.Args[patternIndex]
	paths := os.Args[patternIndex+1:]

	grep(pattern, paths, recursive)
}
