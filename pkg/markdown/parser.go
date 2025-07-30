package markdown

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func FindMarkdownFiles(dirs []string) ([]string, error) {
	var files []string

	for _, dir := range dirs {
		err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if !d.IsDir() && strings.HasSuffix(path, ".md") {
				files = append(files, path)
			}
			return nil
		})

		if err != nil {
			fmt.Println(err)
		}
	}

	return files, nil
}

func ParseMarkdownFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
