package modules

import (
	"fmt"
	"math"
	"os"

	"go.uber.org/atomic"
)

// VCOParams for a given oscillator
type VCOParams struct {
	Freq     *atomic.Float64
	Amp      *atomic.Float64
	Waveform *atomic.String
	Shutdown chan os.Signal
}

// VCO - Voltage Controlled Oscillator
func VCO(sr *SampleRate, shutdown chan os.Signal) (*VCOParams, Output) {
	params := &VCOParams{
		Freq:     atomic.NewFloat64(0),
		Amp:      atomic.NewFloat64(0),
		Waveform: atomic.NewString("sine"),
		Shutdown: shutdown,
	}
	t := 0.0
	var gen func(*float64) float64
	origFreq := params.Freq.Load()
	return params, func(out chan<- Signal) {
		for {
			select {
			case <-shutdown:
				fmt.Println("shutting down oscillator")
				return
			default:
				switch params.Waveform.Load() {
				case "sine":
					gen = func(pt *float64) float64 {
						return math.Sin(2.0 * math.Pi * float64(params.Freq.Load()) * *pt)
					}
				case "square":
					gen = func(pt *float64) float64 {
						return 2*(2*math.Floor(params.Freq.Load()**pt)-math.Floor(2*params.Freq.Load()**pt)) + 1
					}
				case "triangle":
					gen = func(pt *float64) float64 {
						p := 1 / params.Freq.Load()
						return 2*math.Abs(2*((*pt/p)-math.Floor((*pt/p)+0.5))) - 1
					}
				}
				if origFreq != params.Freq.Load() {
					t = (origFreq * t) / params.Freq.Load()
					origFreq = params.Freq.Load()
				}
				y := gen(&t) * params.Amp.Load()
				out <- Signal{y, y}
				t += sr.D(1).Seconds()
			}
		}
	}
}
