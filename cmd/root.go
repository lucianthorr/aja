package cmd

import (
	"fmt"

	"github.com/lucianthorr/aja/midi"
	"github.com/lucianthorr/aja/modules"
	"github.com/spf13/cobra"
)

var midiInput *midi.Interface
var speakerOut chan modules.Signal

func init() {
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(connectCmd)
	rootCmd.AddCommand(disconnectCmd)
}

// Execute the root command for Cobra CLI
func Execute(in *midi.Interface, out chan modules.Signal) *cobra.Command {
	midiInput = in
	speakerOut = out
	return rootCmd
}

var rootCmd = &cobra.Command{
	Use:   "synth",
	Short: "A simple midi capable synthesizer",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("you are here")
	},
}
