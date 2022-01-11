package main

import (
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	readmeFilePath = "./README.md"
	problemsTitle  = "## Problems"
)

func main() {
	fmt.Printf("README  ... ")
	UpdateReadMe()
	fmt.Printf("updated 👌\n")
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

func getSolutionsList() (names []string, err error) {
	err = filepath.WalkDir("./solutions", func(path string, d fs.DirEntry, err error) error {
		if !strings.Contains(path, "test") {
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
