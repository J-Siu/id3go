package cmd

import (
	"fmt"

	"github.com/J-Siu/id3go/helper"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:                   "get <files>",
	DisableFlagsInUseLine: true,
	Short:                 "Get tags of files",
	Long:                  `Get tags of files`,
	Run: func(cmd *cobra.Command, args []string) {

		// Loop through file list
		for j := 0; j < len(args); j++ {
			getTags(cmd, &args[j])
		}

	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}

func getTags(cmd *cobra.Command, path *string) {
	fmt.Println("File :", *path)

	fh := helper.Open(*path)
	if fh == nil {
		return
	}

	for i := 0; i < len(helper.Tags); i++ {
		tagLongName := &helper.Tags[i].Ln
		tagDisplayName := &helper.Tags[i].Dn
		tagVal := fh.Get(&helper.Tags[i])

		// If track / year is 0, assume empty
		if (*tagLongName == "year" || *tagLongName == "track") && tagVal == "0" {
			tagVal = ""
		}

		if tagVal != "" {
			fmt.Println(*tagDisplayName, ":", tagVal)
		}
	}

	fh.Close()
}
