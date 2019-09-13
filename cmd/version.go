package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const ver = "1.0"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Long:  `Print version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("id3go version :", ver)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
