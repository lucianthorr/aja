package modules

import (
	"github.com/lucianthorr/modular_sounds/midi"
	"github.com/uber-go/atomic"
)

type Controller func(*atomic.Float64)

// Note from midi to frequency in Hertz
type Note struct {
	Key  uint8
	Name string
	Freq float64
}

var notes = []Note{
	{Key: 24, Name: "C0", Freq: 16.35},
	{Key: 25, Name: "Csharp0", Freq: 17.32},
	{Key: 26, Name: "D0", Freq: 18.35},
	{Key: 27, Name: "Dsharp0", Freq: 19.45},
	{Key: 28, Name: "E0", Freq: 20.60},
	{Key: 29, Name: "F0", Freq: 21.83},
	{Key: 30, Name: "Fsharp0", Freq: 23.12},
	{Key: 31, Name: "G0", Freq: 24.50},
	{Key: 32, Name: "Gsharp0", Freq: 25.96},
	{Key: 33, Name: "A0", Freq: 27.50},
	{Key: 34, Name: "Asharp0", Freq: 29.14},
	{Key: 35, Name: "B0", Freq: 30.87},
	{Key: 36, Name: "C1", Freq: 32.70},
	{Key: 37, Name: "Csharp1", Freq: 34.65},
	{Key: 38, Name: "D1", Freq: 36.71},
	{Key: 39, Name: "Dsharp1", Freq: 38.89},
	{Key: 40, Name: "E1", Freq: 41.20},
	{Key: 41, Name: "F1", Freq: 43.65},
	{Key: 42, Name: "Fsharp1", Freq: 46.25},
	{Key: 43, Name: "G1", Freq: 49.00},
	{Key: 44, Name: "Gsharp1", Freq: 51.91},
	{Key: 45, Name: "A1", Freq: 55.00},
	{Key: 46, Name: "Asharp1", Freq: 58.27},
	{Key: 47, Name: "B1", Freq: 61.74},
	{Key: 48, Name: "C2", Freq: 65.41},
	{Key: 49, Name: "Csharp2", Freq: 69.30},
	{Key: 50, Name: "D2", Freq: 73.42},
	{Key: 51, Name: "Dsharp2", Freq: 77.78},
	{Key: 52, Name: "E2", Freq: 82.41},
	{Key: 53, Name: "F2", Freq: 87.31},
	{Key: 54, Name: "Fsharp2", Freq: 92.50},
	{Key: 55, Name: "G2", Freq: 98.00},
	{Key: 56, Name: "Gsharp2", Freq: 103.83},
	{Key: 57, Name: "A2", Freq: 110.00},
	{Key: 58, Name: "Asharp2", Freq: 116.54},
	{Key: 59, Name: "B2", Freq: 123.47},
	{Key: 60, Name: "C3", Freq: 130.81},
	{Key: 61, Name: "Csharp3", Freq: 138.59},
	{Key: 62, Name: "D3", Freq: 146.83},
	{Key: 63, Name: "Dsharp3", Freq: 155.56},
	{Key: 64, Name: "E3", Freq: 164.81},
	{Key: 65, Name: "F3", Freq: 174.61},
	{Key: 66, Name: "Fsharp3", Freq: 185.00},
	{Key: 67, Name: "G3", Freq: 196.00},
	{Key: 68, Name: "Gsharp3", Freq: 207.65},
	{Key: 69, Name: "A3", Freq: 220.00},
	{Key: 70, Name: "Asharp3", Freq: 233.08},
	{Key: 71, Name: "B3", Freq: 246.94},
	{Key: 72, Name: "C4", Freq: 261.63},
	{Key: 73, Name: "Csharp4", Freq: 277.18},
	{Key: 74, Name: "D4", Freq: 293.66},
	{Key: 75, Name: "Dsharp4", Freq: 311.13},
	{Key: 76, Name: "E4", Freq: 329.63},
	{Key: 77, Name: "F4", Freq: 349.23},
	{Key: 78, Name: "Fsharp4", Freq: 369.99},
	{Key: 79, Name: "G4", Freq: 392.00},
	{Key: 80, Name: "Gsharp4", Freq: 415.30},
	{Key: 81, Name: "A4", Freq: 440.00},
	{Key: 82, Name: "Asharp4", Freq: 466.16},
	{Key: 83, Name: "B4", Freq: 493.88},
	{Key: 84, Name: "C5", Freq: 523.25},
	{Key: 85, Name: "Csharp5", Freq: 554.37},
	{Key: 86, Name: "D5", Freq: 587.33},
	{Key: 87, Name: "Dsharp5", Freq: 622.25},
	{Key: 88, Name: "E5", Freq: 659.25},
	{Key: 89, Name: "F5", Freq: 698.46},
	{Key: 90, Name: "Fsharp5", Freq: 739.99},
	{Key: 91, Name: "G5", Freq: 783.99},
	{Key: 92, Name: "Gsharp5", Freq: 830.61},
	{Key: 93, Name: "A5", Freq: 880.00},
	{Key: 94, Name: "Asharp5", Freq: 932.33},
	{Key: 95, Name: "B5", Freq: 987.77},
	{Key: 96, Name: "C6", Freq: 1046.50},
	{Key: 97, Name: "Csharp6", Freq: 1108.73},
	{Key: 98, Name: "D6", Freq: 1174.66},
	{Key: 99, Name: "Dsharp6", Freq: 1244.51},
	{Key: 100, Name: "E6", Freq: 1318.51},
	{Key: 101, Name: "F6", Freq: 1396.91},
	{Key: 102, Name: "Fsharp6", Freq: 1479.98},
	{Key: 103, Name: "G6", Freq: 1567.98},
	{Key: 104, Name: "Gsharp6", Freq: 1661.22},
	{Key: 105, Name: "A6", Freq: 1760.00},
	{Key: 106, Name: "Asharp6", Freq: 1864.66},
	{Key: 107, Name: "B6", Freq: 1975.53},
	{Key: 108, Name: "C7", Freq: 2093.00},
	{Key: 109, Name: "Csharp7", Freq: 2217.46},
	{Key: 110, Name: "D7", Freq: 2349.32},
	{Key: 111, Name: "Dsharp7", Freq: 2489.02},
	{Key: 112, Name: "E7", Freq: 2637.02},
	{Key: 113, Name: "F7", Freq: 2793.83},
	{Key: 114, Name: "Fsharp7", Freq: 2959.96},
	{Key: 115, Name: "G7", Freq: 3135.96},
	{Key: 116, Name: "Gsharp7", Freq: 3322.44},
	{Key: 117, Name: "A7", Freq: 3520.00},
	{Key: 118, Name: "Asharp7", Freq: 3729.31},
	{Key: 119, Name: "B7", Freq: 3951.07},
	{Key: 120, Name: "C8", Freq: 4186.01},
	{Key: 121, Name: "Csharp8", Freq: 4434.92},
	{Key: 122, Name: "D8", Freq: 4698.63},
	{Key: 123, Name: "Dsharp8", Freq: 4978.03},
	{Key: 124, Name: "E8", Freq: 5274.04},
	{Key: 125, Name: "F8", Freq: 5587.65},
	{Key: 126, Name: "Fsharp8", Freq: 5919.91},
	{Key: 127, Name: "G8", Freq: 6271.93},
	{Key: 128, Name: "Gsharp8", Freq: 6644.88},
	{Key: 129, Name: "A8", Freq: 7040.00},
	{Key: 130, Name: "Asharp8", Freq: 7458.62},
	{Key: 131, Name: "B8", Freq: 7902.13},
}

// Midi2CV translator
func Midi2CV(m *midi.Interface) Controller {
	noteMap := make(map[uint8]float64)
	for _, note := range notes {
		noteMap[note.Key] = note.Freq
	}
	return func(parameter *atomic.Float64) {
		for {
			select {
			case msg := <-m.Messages:
				if msg.Value == "NoteOn" {
					f := noteMap[msg.Key]
					parameter.Store(f)
				}
			default:
			}

		}
	}
}
