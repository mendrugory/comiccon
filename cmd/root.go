package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "comiccon",
	Short: "comiccon is a CLI tool to download e-comics",
	Long: `A toy project which will help you to download comics and keep them updated from https://the-eye.eu/public/Comics/
building the same folder structure.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
