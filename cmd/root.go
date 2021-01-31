package cmd

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lucianthorr/aja/midi"
	"github.com/lucianthorr/aja/modules"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var midiInput = midi.NewInterface()
var speakerOut chan modules.Signal

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "./config/test.yaml", "path to config file")
	rootCmd.AddCommand(midiCmd)
	midiCmd.AddCommand(listCmd)
}

func Execute() error {
	return rootCmd.Execute()
}

var rootCmd = &cobra.Command{
	Use:   "synth",
	Short: "A simple midi capable synthesizer",
	Run: func(cmd *cobra.Command, args []string) {
		sigs := make(chan os.Signal, 1)

		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

		config, err := LoadConfig()
		if err != nil {
			log.Fatal(err)
		}
		sr := modules.SampleRate(48000)
		bufferSize := sr.N(22 * time.Millisecond)
		speakerOut, err = modules.Init(sr, bufferSize)
		if err != nil {
			log.Fatal(err)
		}
		midi2CV := modules.Midi2CV(midiInput)
		midiInput.Connect(midiInput.GetIndex(config.MidiInput))
		defer midiInput.Disconnect()

		params, vco := modules.VCO(&sr, sigs)
		go midi2CV(params.Freq)
		params.Waveform.Store("sine")
		params.Freq.Store(440.0)
		params.Amp.Store(0.5)
		vco(speakerOut)
	},
}

func initConfig() {
	viper.SetConfigFile(cfgFile)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
