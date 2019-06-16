package modules

import (
	"time"

	"github.com/hajimehoshi/oto"
	"github.com/pkg/errors"
)

// Output is a function that puts a signal onto a channel
type Output func(chan<- Signal)

// SampleRate is the number of samples per second.
// Taken from "github.com/faiface/beep"
type SampleRate int

// D returns the duration of n samples.
func (sr SampleRate) D(n int) time.Duration {
	return time.Second * time.Duration(n) / time.Duration(sr)
}

// N returns the number of samples that last for d duration.
func (sr SampleRate) N(d time.Duration) int {
	return int(d * time.Duration(sr) / time.Second)
}

// Signal for stereo audio.  0: Left, 1: Right
type Signal [2]float64

// Add two signals
func (s *Signal) Add(t *Signal) *Signal {
	return &Signal{s[0] + t[0], s[1] + t[1]}
}

// Multiply two signals (assuming their the same length)
func (s *Signal) Multiply(t *Signal) *Signal {
	return &Signal{s[0] * t[0], s[1] * t[1]}
}

var (
	speakerIn chan Signal
	buf       []byte
	player    *oto.Player
	done      chan struct{}
)

// Init the system to play out.
// Also basically taken from "github.com/faiface/beep"
func Init(sr SampleRate, bufferSize int) (chan Signal, error) {

	if player != nil {
		done <- struct{}{}
		player.Close()
	}

	numBytes := bufferSize * 4
	speakerIn = make(chan Signal, bufferSize)
	buf = make([]byte, numBytes)

	var err error
	player, err = oto.NewPlayer(int(sr), 2, 2, numBytes)
	if err != nil {
		return speakerIn, errors.Wrap(err, "failed to initialize speaker")
	}
	done = make(chan struct{})

	go func() {
		for {
			select {
			default:
				update(bufferSize)
			case <-done:
				return
			}
		}
	}()
	return speakerIn, nil
}

func update(bufferSize int) {
	i := 0
	for i < bufferSize {
		signal := <-speakerIn
		for c, val := range signal {
			if val < -1 {
				val = -1
			}
			if val > 1 {
				val = 1
			}
			valInt16 := int16(val * (1<<15 - 1))
			low := byte(valInt16)
			high := byte(valInt16 >> 8)
			buf[i*4+c*2+0] = low
			buf[i*4+c*2+1] = high
		}
		i += 2
	}
	player.Write(buf)
}
