package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var midiCmd = &cobra.Command{
	Use:   "midi",
	Short: "Signifies MIDI related commands",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This command requires subcommands")
		subCmds := cmd.Commands()
		for i, subCmd := range subCmds {
			fmt.Printf("%d:\t%s\n", i, subCmd.Use)
		}
	},
}
var listCmd = &cobra.Command{
	Use:   "ls",
	Short: "List Available Midi Devices",
	Run: func(cmd *cobra.Command, args []string) {
		midiInput.List()
	},
}
