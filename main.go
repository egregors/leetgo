// main is main and linters are linters ¯\_(ツ)_/¯
package main

import (
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"os/exec"
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

	fmt.Printf("💬 COMMENTS ... \n")
	NormalizeComments(false)
	fmt.Printf("👌 normalized \n")

	fmt.Printf("✅ NEW SOLUTIONS ... \n")
	AddToGit()
	fmt.Printf("👌 git added \n")

	fmt.Printf("📚 README  ... \n")
	UpdateReadMe()
	fmt.Printf("👌 updated \n")
}

func AddToGit() {
	cmd := "git add ./solutions"
	out, err := execCommand(cmd)
	if err != nil {
		log.Fatalf("git add failed: %s, output: %s", err, out)
	}

	cmd = "git add ./README.md"
	out, err = execCommand(cmd)
	if err != nil {
		log.Fatalf("git add README failed: %s, output: %s", err, out)
	}
}

func execCommand(cmd string) (string, error) {
	out, err := exec.Command("bash", "-c", cmd).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("command failed: %s, output: %s", err, out)
	}
	return strings.TrimSpace(string(out)), nil
}

// NormalizeNames renames solutions files to match the pattern s0001_problem_name.go
func NormalizeNames(dry bool) {
	err := filepath.WalkDir("./solutions", func(path string, d fs.DirEntry, _ error) error {
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
	err = filepath.WalkDir("./solutions", func(path string, _ fs.DirEntry, _ error) error {
		if isValid(path) {
			words := strings.Split(path, "/")
			problemName := strings.Split(words[1], ".")[0]
			problemName = "* [" + strings.ReplaceAll(problemName, "_", " ") + "](" + path + ")"
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

// NormalizeComments formats comments in solution files to follow proper formatting rules:
// - lines should not exceed 80 characters
// - word breaks only at whitespace or other break characters
// - entire comment block should be indented at least one tab from left edge
func NormalizeComments(dry bool) {
	err := filepath.WalkDir("./solutions", func(path string, d fs.DirEntry, _ error) error {
		skip := map[string]bool{
			"main_test.go": true,
			"types.go":     true,
			"utils.go":     true,
		}

		if !d.IsDir() && !skip[d.Name()] && strings.HasSuffix(d.Name(), ".go") && !strings.HasSuffix(d.Name(), "_test.go") {
			err := normalizeFileComments(path, dry)
			if err != nil {
				log.Printf("Error processing %s: %v", path, err)
			}
		}

		return nil
	})
	if err != nil {
		log.Fatalf("comment normalization failed: %s", err)
	}
}

func normalizeFileComments(filePath string, dry bool) error {
	// Validate file path to prevent path traversal attacks
	cleanPath := filepath.Clean(filePath)
	if !strings.HasPrefix(cleanPath, "solutions/") && !strings.HasPrefix(cleanPath, "./solutions/") {
		return fmt.Errorf("invalid file path: %s", filePath)
	}

	content, err := os.ReadFile(cleanPath) // #nosec G304 - path is validated above
	if err != nil {
		return err
	}

	lines := strings.Split(string(content), "\n")
	if len(lines) < 2 {
		return nil
	}

	// Find the comment block (starts with /* and ends with */)
	startIdx := -1
	endIdx := -1
	for i, line := range lines {
		if strings.Contains(line, "/*") {
			startIdx = i
		}
		if strings.Contains(line, "*/") {
			endIdx = i
			break
		}
	}

	if startIdx == -1 || endIdx == -1 || startIdx >= endIdx {
		return nil // No comment block found
	}

	// Process each line of the comment block individually
	commentLines := lines[startIdx : endIdx+1]
	var formattedLines []string

	for i, line := range commentLines {
		if i == 0 {
			// First line with /*
			if idx := strings.Index(line, "/*"); idx != -1 {
				prefix := line[:idx]
				after := line[idx+2:]
				if strings.TrimSpace(after) == "" {
					formattedLines = append(formattedLines, prefix+"/*")
				} else {
					// Check if the line is too long and wrap if needed
					wrapped := wrapCommentLine(prefix+"/*"+after, 80)
					formattedLines = append(formattedLines, wrapped...)
				}
			}
		} else if i == len(commentLines)-1 {
			// Last line with */
			if idx := strings.Index(line, "*/"); idx != -1 {
				before := line[:idx]
				suffix := line[idx:]
				if strings.TrimSpace(before) == "" {
					formattedLines = append(formattedLines, before+suffix)
				} else {
					// Check if the line is too long and wrap if needed
					wrapped := wrapCommentLine(before+suffix, 80)
					formattedLines = append(formattedLines, wrapped...)
				}
			}
		} else {
			// Middle lines - preserve structure but wrap long lines
			if strings.TrimSpace(line) == "" {
				// Preserve empty lines
				formattedLines = append(formattedLines, line)
			} else {
				// Wrap long lines while preserving indentation
				wrapped := wrapCommentLine(line, 80)
				formattedLines = append(formattedLines, wrapped...)
			}
		}
	}

	if dry {
		fmt.Printf("File: %s\n", filePath)
		for _, line := range formattedLines {
			fmt.Println(line)
		}
		fmt.Println(strings.Repeat("-", 80))
		return nil
	}

	// Replace the comment block
	newLines := make([]string, 0, len(lines))
	newLines = append(newLines, lines[:startIdx]...)
	newLines = append(newLines, formattedLines...)
	newLines = append(newLines, lines[endIdx+1:]...)

	newContent := strings.Join(newLines, "\n")
	return os.WriteFile(filePath, []byte(newContent), 0o600)
}

// wrapCommentLine wraps a single comment line if it exceeds maxLength
// while preserving the original indentation
func wrapCommentLine(line string, maxLength int) []string {
	if len(line) <= maxLength {
		return []string{line}
	}

	// Extract leading whitespace/indentation
	leadingSpace := ""
	for i, r := range line {
		if r != ' ' && r != '\t' {
			leadingSpace = line[:i]
			break
		}
	}

	// Get the actual content without leading space
	content := strings.TrimLeft(line, " \t")

	// If it's still not too long, return as-is
	if len(leadingSpace+content) <= maxLength {
		return []string{leadingSpace + content}
	}

	var result []string
	words := strings.Fields(content)
	if len(words) == 0 {
		return []string{line}
	}

	currentLine := leadingSpace + words[0]

	for i := 1; i < len(words); i++ {
		word := words[i]
		testLine := currentLine + " " + word

		if len(testLine) > maxLength {
			// Current line is full, start a new one
			result = append(result, currentLine)
			currentLine = leadingSpace + "\t" + word
		} else {
			currentLine = testLine
		}
	}

	// Add the last line
	if currentLine != leadingSpace {
		result = append(result, currentLine)
	}

	return result
}
