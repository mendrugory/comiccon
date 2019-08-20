package worker

import (
	"testing"
)

// func NewLink(url string, suffixes []string) {

// }

func TestIsResource(t *testing.T) {
	url := "https://www.myweb.com/myfile.pdf"
	suffixes := []string{"pdf", "jpg"}
	if !isResource(url, suffixes) {
		t.Fatalf("The url: %s does not belong to a resource", url)
	}
}

