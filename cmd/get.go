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
	"github.com/J-Siu/id3go/tag"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:                   "get <files>",
	Aliases:               []string{"g"},
	DisableFlagsInUseLine: true,
	Short:                 "Get tags of files",
	Run: func(cmd *cobra.Command, args []string) {
		// Loop through file list
		for j := 0; j < len(args); j++ {
			getTags(&args[j])
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}

func getTags(path *string) {
	ezlog.Log().N("File").M(path).Out()

	t := new(tag.TagFile).New(*path)
	if t.Err != nil {
		return
	}

	for _, item := range tag.Tags {
		tagLongName := &item.Ln
		tagDisplayName := &item.Dn
		tagVal := t.Get(&item)

		// If track / year is 0, assume empty
		if (*tagLongName == "year" || *tagLongName == "track") && tagVal == "0" {
			tagVal = ""
		}
		if tagVal != "" {
			ezlog.Log().N(tagDisplayName).M(tagVal).Out()
		}
	}

	t.Close()
}
