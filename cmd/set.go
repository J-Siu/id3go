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
