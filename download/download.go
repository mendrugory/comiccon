package download

import (
	"net/http"
	"io/ioutil"
)

// Download downloads the file and returns the received bytes and the error.
func Download(url string) ([]byte, error) {
	resp, err := http.Get(url)
    if err != nil {
        return nil, err
	}
	
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
        return nil, err
	}
	
	return data, nil
}
