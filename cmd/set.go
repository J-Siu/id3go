/*
The MIT License

Copyright (c) 2020 John Siu

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package cmd

import (
	"fmt"

	"github.com/J-Siu/id3go/helper"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set [flags] <files>",
	Short: "Set tags of files",
	Long:  `Set tags of files`,
	Run: func(cmd *cobra.Command, args []string) {

		save, _ := cmd.Flags().GetBool("Save")
		dryrunMsg := ""
		if !save {
			dryrunMsg = "(Dry Run)"
		}

		// Check if any flag is used
		anyChange := false
		for i := 0; i < len(helper.Tags); i++ {
			anyChange = (anyChange || cmd.Flag(helper.Tags[i].Ln).Changed)
		}

		// Return if no flag is used (no tag to update)
		if !anyChange {
			fmt.Println("Nothing to update.")
			return
		}

		// Loop through file list
		for j := 0; j < len(args); j++ {
			setTags(cmd, &args[j], &dryrunMsg, save)
		}
	},
}

func init() {
	rootCmd.AddCommand(setCmd)

	// Loop through Use helper.Tags to setup all flags
	for i := 0; i < len(helper.Tags); i++ {
		tagLongName := &helper.Tags[i].Ln
		tagShortName := &helper.Tags[i].Sn
		tagMessage := &helper.Tags[i].Ms

		setCmd.Flags().StringP(*tagLongName, *tagShortName, "", *tagMessage)
	}
	// "Save" flag
	setCmd.Flags().BoolP("Save", "S", false, "save to file. Without this flag, `set` will not writing to files (dry run).")
}

func setTags(cmd *cobra.Command, path *string, dryrunMsg *string, save bool) {

	fmt.Println("File", *dryrunMsg, ":", *path)

	fh := helper.Open(*path)
	if fh == nil {
		// Fail to open, next file ...
		return
	}

	updateMsg := ""
	for i := 0; i < len(helper.Tags); i++ {
		flagLongName := &helper.Tags[i].Ln
		flagChanged := cmd.Flag(*flagLongName).Changed
		tagDisplayName := &helper.Tags[i].Dn
		tagValNew := cmd.Flag(*flagLongName).Value.String()
		tagValOld := fh.Get(&helper.Tags[i])
		if flagChanged && tagValNew != tagValOld {
			updateMsg += fmt.Sprintln(*tagDisplayName, ":", tagValOld, "->", tagValNew)
			fh.Set(&helper.Tags[i], tagValNew)
		}
	}
	if updateMsg == "" {
		// All new tag values are same as old
		fmt.Println("Nothing to update.")
	} else {

		fmt.Print(updateMsg)

		// Only save to file if not dry run
		if save {
			fh.Save()
			fmt.Println("Updated.")
		}
	}

	fh.Close()
}
