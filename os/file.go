package os

import (
	"io/ioutil"
	"os"
)

const fileMode os.FileMode = 0644

// SaveBinaryFile saves the given binary content in the given filePath. It returns the error.
func SaveBinaryFile(binary []byte, filePath string) error {
	return ioutil.WriteFile(filePath, binary, fileMode)
}

// ReadBinaryFile returns the content of the file
func ReadBinaryFile(filePath string) ([]byte, error) {
	return ioutil.ReadFile(filePath)
}
