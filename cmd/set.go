/*
The MIT License

Copyright (c) 2025 John Siu

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package cmd

import (
	"github.com/J-Siu/go-helper/v2/ezlog"
	"github.com/J-Siu/id3go/global"
	"github.com/J-Siu/id3go/tag"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:     "set [flags] <files>",
	Aliases: []string{"s"},
	Short:   "Set tags of files",
	Long:    `Set tags of files`,
	Run: func(cmd *cobra.Command, args []string) {
		save, _ := cmd.Flags().GetBool("Save")
		dryrunMsg := ""
		if !save {
			dryrunMsg = "(Dry Run)"
		}
		for _, item := range args {
			setTags(cmd, &item, &dryrunMsg, save)
		}
	},
}

func init() {
	cmd := setCmd
	rootCmd.AddCommand(cmd)
	// Loop through Use tag.Tags to setup all flags
	for _, item := range global.Tags {
		tagLongName := &item.Ln
		tagShortName := &item.Sn
		tagMessage := &item.Ms
		cmd.Flags().StringP(*tagLongName, *tagShortName, "", *tagMessage)
	}
	// "Save" flag
	cmd.Flags().BoolP("Save", "S", false, "save to file. Without this flag, `set` will not update file (dry run).")
}

func setTags(cmd *cobra.Command, path, dryrunMsg *string, save bool) {
	ezlog.Log().N("File" + *dryrunMsg).M(path).Out()

	t := new(tag.TagFile).New(*path)
	if t.Err != nil {
		return // Fail to open, next file ...
	}

	updated := false
	ezlog.Log()
	for _, item := range global.Tags {
		flagChanged := cmd.Flag(item.Ln).Changed
		tagDisplayName := &item.Dn
		tagValNew := cmd.Flag(item.Ln).Value.String()
		tagValOld := t.Get(&item)
		if flagChanged && tagValNew != tagValOld {
			t.Set(&item, tagValNew)
			updated = true
			ezlog.N(tagDisplayName).M(tagValOld).M("->").Mn(tagValNew)
		}
	}
	if updated {
		ezlog.Se().Out()
		// Only save to file if not dry run
		if save {
			t.Save()
			ezlog.M("Updated.").Out()
		}
	} else {
		// All new tag values are same as old
		ezlog.M("Nothing to update.").Out()
	}

	t.Close()
}
