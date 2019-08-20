package cmd

import (
	"fmt"
	"net/url"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/mendrugory/comicon/data"
	"github.com/mendrugory/comicon/manager"
	helper "github.com/mendrugory/comicon/os"
	"github.com/spf13/cobra"
)

const baseFolderFlag string = "basefolder"
const suffixesFlag string = "suffixes"
const linkFlag string = "link"

const baseUrl string = "https://the-eye.eu/public/Comics"
const defaultBaseFolder string = "./comics"
const defaultSuffixes string = "cbr,jpg,pdf"
const defaultLink string = ""

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download comics",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Downloading comics")

		baseFolder := getBaseFolder(cmd, baseFolderFlag)
		fmt.Println("Basefolder:", baseFolder)

		url := getURL(cmd, linkFlag)
		fmt.Println("Url:", url)

		suffixes := getSuffixes(cmd, suffixesFlag)
		fmt.Println("Suffixes:", suffixes)

		if result := run(url, baseFolder, suffixes); !result {
			fmt.Println("Errors during downloading 😱")
			os.Exit(1)
		}

		fmt.Println("Check your comics 📖")
	},
}

func getFlagValue(cmd *cobra.Command, flagName string, defaultValue string) string {
	if value, err := cmd.Flags().GetString(flagName); err != nil {
		fmt.Printf("Error with flag %s: %v\n", flagName, err)
		return defaultValue
	} else if value == "" {
		return defaultValue
	} else {
		return value
	}
}

func getBaseFolder(cmd *cobra.Command, flagName string) string {
	return getFlagValue(cmd, flagName, defaultBaseFolder)
}

func getURL(cmd *cobra.Command, flagName string) string {
	link := getFlagValue(cmd, flagName, defaultLink)
	u, _ := url.Parse(baseUrl)
	u.Path = path.Join(u.Path, link)

	result, err := url.QueryUnescape(u.String())
	if err != nil {
		fmt.Printf("Error with flag %s: %v\n", flagName, err)
		return baseUrl
	}

	return result
}

func getSuffixes(cmd *cobra.Command, flagName string) []string {

	stringSuffixes := getFlagValue(cmd, flagName, defaultSuffixes)
	var suffixes []string

	for _, s := range strings.Split(stringSuffixes, ",") {
		suffixes = append(suffixes, s)
	}

	return suffixes
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	downloadCmd.Flags().String(baseFolderFlag, defaultBaseFolder, "Base Folder where the comics will be stored")
	downloadCmd.Flags().String(suffixesFlag, defaultSuffixes, "Suffixes of the files which will be downloaded")
	downloadCmd.Flags().String(linkFlag, defaultLink, "Sublink of the url if you only want to download some collections")
}

func run(url string, baseFolder string, suffixes []string) bool {
	d := data.Resource{
		Url:        url,
		Suffixes:   suffixes,
		BaseFolder: baseFolder,
	}

	helper.CreateFolder(baseFolder)
	d.SaveToFile()

	cpus := runtime.NumCPU()
	fmt.Println("Number of CPUs:", cpus)
	c := manager.Run(cpus, d)

	return <-c
}
