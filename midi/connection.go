package midi

import (
	"fmt"
	"io"
	"log"
	"time"

	"github.com/gomidi/rtmididrv"
	"gitlab.com/gomidi/midi/mid"
)

// Interface for system to communicate with
type Interface struct {
	connected bool
	Messages  chan Message
	reader    *mid.Reader
}

// NewInterface initialized
func NewInterface() *Interface {
	return &Interface{
		connected: false,
		Messages:  make(chan Message),
	}
}

// Connect to the midi device
func (m *Interface) Connect(index int) {
	drv, err := rtmididrv.New()
	if err != nil {
		log.Println(err)
	}
	defer drv.Close()
	ins, err := drv.Ins()
	if err != nil {
		log.Println(err)
	}
	portNum := index - 1
	in := ins[portNum]
	if err := in.Open(); err != nil {
		log.Println(err)
	}

	if err := m.configureReader(in); err != nil {
		log.Println(err)
		return
	}
	// TODO: change shared variables to use sync package and Waits vs loops.
	for m.connected == true {
		time.Sleep(time.Millisecond)
		// select {
		// case msg := <-m.messages:
		// 	fmt.Println(msg)
		// default:
		// }
	}
}

func (m *Interface) configureReader(in mid.In) error {
	m.reader = mid.NewReader(mid.NoLogger())
	mid.ConnectIn(in, m.reader)
	m.connected = true
	m.reader.Msg.Channel.NoteOn = m.enqueueMessage("NoteOn")
	m.reader.Msg.Channel.NoteOff = m.enqueueMessage("NoteOff")
	err := m.reader.ReadAll()
	if err != nil && err != io.EOF {
		return err
	}
	return nil
}

// enqueueMessages coming from midi into messages that we care about
func (m *Interface) enqueueMessage(val string) func(*mid.Position, uint8, uint8, uint8) {
	return func(p *mid.Position, c, k, v uint8) {
		m.Messages <- Message{
			Value:    val,
			Channel:  c,
			Key:      k,
			Velocity: v,
		}
	}
}

// Disconnect from midi port
func (m *Interface) Disconnect() {
	m.connected = false
	fmt.Println("Disconnected from device")
}

// List the available MIDI devices
func (m *Interface) List() {
	drv, err := rtmididrv.New()
	if err != nil {
		log.Println(err)
	}
	defer drv.Close()
	ins, err := drv.Ins()
	if err != nil {
		log.Println(err)
	}
	for i := range ins {
		fmt.Printf("%d:\t%s\n", i+1, ins[i].String())
	}
}
