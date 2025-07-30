package markdown

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"testing"
)

func TestFindMarkdownFiles(t *testing.T) {
	// Create a temporary directory
	tmpDir, err := os.MkdirTemp("", "gopilot-test")
	require.NoError(t, err)

	// Defer the call before the surrounding function returns
	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {

		}
	}(tmpDir)

	// Write some markdown files
	err = os.WriteFile(tmpDir+"/test.md", []byte("# Test 1"), 0644)
	require.NoError(t, err)

	// Write a text file that should not be parsed
	err = os.WriteFile(tmpDir+"/not-markdown.txt", []byte("test2"), 0644)
	require.NoError(t, err)

	// Test subdirectories
	subDir := filepath.Join(tmpDir, "subdir")
	err = os.Mkdir(subDir, 0755)
	require.NoError(t, err)

	err = os.WriteFile(subDir+"/test.md", []byte("# Test 2"), 0644)
	require.NoError(t, err)

	// Run the function and assert the results
	var dirs = []string{tmpDir}
	files, err := FindMarkdownFiles(dirs)
	require.NoError(t, err)
	require.Len(t, files, 2)
	expectedFiles := []string{
		filepath.Join(tmpDir, "test.md"),
		filepath.Join(subDir, "test.md"),
	}
	assert.ElementsMatchf(t, expectedFiles, files, "Expected files: %v, got: %v", expectedFiles, files)
}

func TestParseMarkdownFile(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "gopilot-test")
	require.NoError(t, err)

	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {
		}
	}(tmpDir)

	filePath := filepath.Join(tmpDir, "test.md")
	fileContent := "# Test 1\n\nThis is a test file"
	err = os.WriteFile(filePath, []byte(fileContent), 0644)
	require.NoError(t, err)

	content, err := ParseMarkdownFile(filePath)
	require.NoError(t, err)
	assert.Equal(t, fileContent, content)

	_, err = ParseMarkdownFile(filePath + "invalid")
	assert.Error(t, err)
}
