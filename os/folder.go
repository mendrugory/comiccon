package os

import (
	"fmt"
	"os"
)

// CreateFolder creates the entire folder structure of directory
func CreateFolder(folderPath string) {
	os.MkdirAll(folderPath, os.ModePerm)
}

// DeleteFolder deletes the given folder if it exists.
func DeleteFolder(folderPath string) {
	fmt.Printf(folderPath)
	if _, err := os.Stat(folderPath); !os.IsNotExist(err) {
		os.RemoveAll(folderPath)
	}
}
