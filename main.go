// main is main and linters are linters ¯\_(ツ)_/¯
package main

import (
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	readmeFilePath = "./README.md"
	problemsTitle  = "## Problems"
)

func main() {
	fmt.Printf("⛳️ SOLUTIONS ... \n")
	NormalizeNames(false)
	fmt.Printf("👌 formated \n")

	fmt.Printf("📚 README  ... \n")
	UpdateReadMe()
	fmt.Printf("👌 updated \n")
}

// NormalizeNames renames solutions files to match the pattern s0001_problem_name.go
func NormalizeNames(dry bool) {
	err := filepath.WalkDir("./solutions", func(path string, d fs.DirEntry, err error) error {
		skip := map[string]bool{
			"main_test.go": true,
			"types.go":     true,
			"utils.go":     true,
		}

		if !d.IsDir() && !skip[d.Name()] && !matchSolutionName(d.Name()) {
			problemName := d.Name()
			problemName = strings.ToLower(problemName)
			problemName = strings.ReplaceAll(problemName, " ", "_")
			problemName = strings.Replace(problemName, ".", "", 1)
			xs := strings.Split(problemName, "_")
			for len(xs[0]) < 4 {
				xs[0] = "0" + xs[0]
			}
			newName := "s" + xs[0] + "_" + strings.Join(xs[1:], "_")
			newPath := "./solutions/" + newName
			if path != newPath {
				fmt.Printf("  %-60s -> %s\n", path, newPath)
				if !dry {
					renameFile(path, newPath)
				}
			}

		}

		return nil
	})
	if err != nil {
		log.Fatalf("format failed: %s", err)
	}
}

func renameFile(oldPath, newPath string) {
	err := os.Rename(oldPath, newPath)
	if err != nil {
		log.Fatalf("can't rename file: %s", err)
	}
}

func matchSolutionName(name string) bool {
	pattern := `^s\d{4}_[\w()]+(\s*-\s*\w+)*\.go$`
	re := regexp.MustCompile(pattern)

	return re.MatchString(name)
}

// UpdateReadMe syncs solutions list in repo README.md file
func UpdateReadMe() {
	file, err := os.OpenFile(readmeFilePath, os.O_RDWR, 0o600)
	if err != nil {
		log.Fatalf("can't open the file: %s", err)
	}

	offset, _ := getOffsetOf(file, []byte(problemsTitle))
	solutions, _ := getSolutionsList()
	data := strings.Join(solutions[1:], "\n")

	_, err = file.WriteAt([]byte("\n\n"+data), offset+int64(len(problemsTitle)))
	if err != nil {
		log.Panicf("can't write file: %s", err)
	}

	err = file.Close()
	if err != nil {
		log.Fatalf("can't close the file: %s", err)
	}
}

func isValid(path string) bool {
	invalidTokens := []string{
		"test.go",
		"utils.go",
		"types.go",
	}

	for _, t := range invalidTokens {
		if strings.Contains(path, t) {
			return false
		}
	}

	return true
}

func getSolutionsList() (names []string, err error) {
	err = filepath.WalkDir("./solutions", func(path string, d fs.DirEntry, err error) error {
		if isValid(path) {
			words := strings.Split(path, "/")
			problemName := strings.Split(words[1], ".")[0]
			problemName = "* [" + strings.Replace(problemName, "_", " ", -1) + "](" + path + ")"
			names = append(names, problemName)
		}

		return nil
	})

	return names, err
}

func getOffsetOf(f io.ReadSeeker, search []byte) (int64, error) {
	data, err := io.ReadAll(f)
	if err != nil {
		return 0, err
	}

	return int64(bytes.Index(data, search)), nil
}
