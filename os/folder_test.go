package os

import (
	"fmt"
	"os"
	"path"
	"testing"
)

// TestCreateFolder is a Test for CreateFolder
func TestCreateFolder(t *testing.T) {
	baseFolder := "/tmp"
	folderName := "gotest"
	directory := path.Join(baseFolder, folderName)
	defer DeleteFolder(directory)
	CreateFolder(directory)
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		t.Fatalf(fmt.Sprintf("Folder %s does not exist", directory))
	}
}
