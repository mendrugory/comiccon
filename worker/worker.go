package worker

import (
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/anaskhan96/soup"
	"github.com/mendrugory/comicon/data"
	"github.com/mendrugory/comicon/download"
	"github.com/mendrugory/comicon/os"
)

var notValidLinks = []string{".", "..", "..."}

// Download manages the given the resource
func Download(d data.Resource, releaseC chan bool, jobsC chan data.Resource) {

	defer func() {
		fmt.Println("Finishing Job related to url:", d.Url)
		releaseC <- true
		fmt.Println("Finished Job related to url:", d.Url)
	}()

	executable := isDownloable(d)

	if executable {
		fmt.Printf("Downloading %s...\n", d.Url)

		dowloadedData, err := download.Download(d.Url)
		d.Data = dowloadedData
		if err != nil {
			fmt.Printf("%s", err)
			return
		}

		if isResource(d.Url, d.Suffixes) {
			saveFile(d)
		} else {
			goDeeper(d, jobsC)
		}

	}
}

func isDownloable(d data.Resource) bool {
	isFile := d.IsFile()
	return d.Url != "" && (!isFile || (isFile && d.Size() < data.MinimumResourceSize))
}

func newURL(baseURL string, suffix string) string {
	u, err := url.Parse(baseURL)
	if err != nil {
		fmt.Printf("Error Parsing %s\n", baseURL)
		return ""
	}
	u.Path = path.Join(u.Path, suffix)
	result, err := url.QueryUnescape(u.String())
	if err != nil {
		fmt.Printf("Error Decoding %s\n", u.String())
		return ""
	}
	return result
}

func goDeeper(d data.Resource, jobsC chan data.Resource) {
	folderPath := d.GetFolderPath()
	os.CreateFolder(folderPath)
	for _, link := range getLinks(d) {
		newData := data.Resource{
			Url:        newURL(d.Url, link),
			Suffixes:   d.Suffixes,
			BaseFolder: folderPath,
		}
		jobsC <- newData
	}
}

func isValidLink(link string) bool {
	for _, l := range notValidLinks {
		if l == link {
			return false
		}
	}
	return true
}

func decode(link string) string {
	decoded, err := url.QueryUnescape(link)
	if err != nil {
		return link
	}
	return decoded
}

func getLinks(d data.Resource) []string {

	var result []string

	download, err := download.Download(d.Url)
	if err != nil {
		fmt.Printf("Error %s with %s", err, d.Url)
	}

	html := soup.HTMLParse(string(download))
	pre := html.Find("pre")
	if pre.Error == nil {
		for _, link := range pre.FindAll("a") {
			href := link.Attrs()["href"]
			if isValidLink(href) {
				result = append(result, decode(href))
			}
		}
	}
	return result
}

func isResource(url string, suffixes []string) bool {
	result := false
	for _, suffix := range suffixes {
		result = result || strings.HasSuffix(url, suffix)
	}
	return result
}

func saveFile(data data.Resource) {
	os.SaveBinaryFile(data.Data, data.GetFilePath())
}
