package cmd

import (
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List Available Midi Devices",
	Run: func(cmd *cobra.Command, args []string) {
		midiInput.List()
	},
}

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to a given device",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			index, err := strconv.Atoi(args[0])
			if err != nil {
				log.Println(err)
			}
			midiInput.Connect(index)
		}
	},
}

var disconnectCmd = &cobra.Command{
	Use:   "disconnect",
	Short: "Disconnects from current midi device",
	Run: func(cmd *cobra.Command, args []string) {
		midiInput.Disconnect()
	},
}
