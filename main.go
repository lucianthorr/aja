package main

import (
	"flag"
	"fmt"
	"log"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"
	"strings"
	"time"

	"github.com/lucianthorr/aja/cmd"
	"github.com/lucianthorr/aja/midi"
	"github.com/lucianthorr/aja/modules"
	"github.com/manifoldco/promptui"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")

var midiInput = midi.NewInterface()
var speakerOut chan modules.Signal

func runAudio() {
	var err error
	sr := modules.SampleRate(48000)
	bufferSize := sr.N(22 * time.Millisecond)
	speakerOut, err = modules.Init(sr, bufferSize)
	if err != nil {
		log.Fatal(err)
	}
	midi2CV := modules.Midi2CV(midiInput)

	params, vco := modules.VCO(&sr)
	go midi2CV(params.Freq)
	params.Waveform.Store("sine")
	params.Freq.Store(440.0)
	params.Amp.Store(0.5)
	vco(speakerOut)
}

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}
	go runAudio()
	rootCmd := cmd.Execute(midiInput, speakerOut)
	prompt := promptui.Prompt{
		Label: ">",
	}
	for {
		result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		splitResult := strings.Split(result, " ")
		for _, command := range rootCmd.Commands() {
			if splitResult[0] == command.Use {
				args := []string{}
				if len(splitResult) > 0 {
					args = splitResult[1:len(splitResult)]
				}
				go command.Run(command, args)
			}
		}
		if result == "q" || result == "quit" || result == "exit" {
			break
		}
	}

}
