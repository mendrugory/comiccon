package worker

import (
	"testing"
)

// func NewLink(url string, extensions []string) {

// }

func TestIsResource(t *testing.T) {
	url := "https://www.myweb.com/myfile.pdf"
	extensions := []string{"pdf", "jpg"}
	if !isResource(url, extensions) {
		t.Fatalf("The url: %s does not belong to a resource", url)
	}
}

