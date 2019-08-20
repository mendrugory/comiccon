package data

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"

	helper "github.com/mendrugory/comicon/os"
)

type Resource struct {
	Url        string   `json:"url"`
	Suffixes   []string `json:"suffixes"`
	BaseFolder string   `json:"basefolder"`
	Data       []byte	`json:"-"`
	filePath   string `json:"filepath"`
}

// MinimumResourceSize is the limit to not downloaded the resource again.
const MinimumResourceSize int64 = 5500 //bytes

const resourceConfigFileName = "config.json"

// GetFilePath returns the file path
func (d *Resource) GetFilePath() string {
	if d.filePath == "" {
		urls := strings.Split(d.Url, "/")
		name := urls[len(urls)-1]
		d.filePath = path.Join(d.BaseFolder, name)
	}
	return d.filePath
}

// GetFolderPath returns the folder path
func (d *Resource) GetFolderPath() string {
	return d.GetFilePath()
}

// IsDownloaded checks if the Resource has been dowloaded
func (d Resource) IsDownloaded() bool {
	filePath := d.GetFilePath()
	_, err := os.Stat(filePath)
	return err == nil
}

// IsFile checks if a Resource is downloaded and it is a file.
func (d Resource) IsFile() bool {
	filePath := d.GetFilePath()
	if f, err := os.Stat(filePath); err == nil {
		return !f.IsDir()
	}
	return false
}

// Size returns the size of the resource. 0 if it does not exist.
func (d Resource) Size() int64 {
	filePath := d.GetFilePath()
	if f, err := os.Stat(filePath); err == nil {
		return f.Size()
	}
	return 0
}

func (d Resource) SaveToFile() {
	filePath := path.Join(d.BaseFolder, resourceConfigFileName)
	b, err := json.Marshal(d)
	if err != nil {
		fmt.Printf("Error Saving configuration: %s", err)
		return
	}

	if helper.SaveBinaryFile(b, filePath) != nil {
		fmt.Printf("Error Saving configuration: %s", err)
	}
}

// func RetrieveConfiguration(filePath) Resource {

// }
