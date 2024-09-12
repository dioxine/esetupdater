/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"esetupdater/internal"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var fullFlag bool = false

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Execute update process. (-f flag for full download)",
	Long: `
	This command runs update process from the server specified in config. 
	If -f or --full flag spefified, then download the whole base of updates, overwriting
	existing files.`,
	Run: func(cmd *cobra.Command, args []string) {
		var fullDownload bool
		// Check if it is the first start and there is no update.old file OR it was given
		// flag -f, toggle fullDownload to true and download whole base of updates
		// accordingly to filter setting in config.

		flag, _ := cmd.Flags().GetBool("full")

		if internal.CheckIfFirstStart() || flag {
			fullDownload = true
		} else {
			fullDownload = false
		}

		fmt.Println("This is full download:", fullDownload)

		// Check if update.ver exists, if yes - rename it to update.old
		if err := internal.RenameUpdateVer(); err != nil {
			log.Fatalln(err)
		}

		// Download fresh update.ver
		_, err := internal.DownloadFile("/dll/update.ver")
		if err != nil {
			log.Fatalln(err)
		}

		// Get old and new update.ver converted to arrays of SectionBlock maps and compare
		// their lenghts. Return old and new arrays.
		sectionBlockArrayOld, sectionBlockArrayNew, err := internal.CompareUpdateVer()
		if err != nil {
			log.Fatalln(err)
		}

		// Run update process (full or delta, depenging on fullDownload boolean)
		internal.UpdateMain(sectionBlockArrayOld, sectionBlockArrayNew, fullDownload)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// -f or --full flag initiation
	updateCmd.Flags().BoolVarP(&fullFlag, "full", "f", false, "Force full update")
}
