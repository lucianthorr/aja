package cmd

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	MidiInput string
}

func LoadConfig() (*Config, error) {
	mi := viper.GetString("midi.input")
	if mi == "" {
		return nil, fmt.Errorf("No Midi Input found in config")
	}
	return &Config{
		MidiInput: mi,
	}, nil
}
