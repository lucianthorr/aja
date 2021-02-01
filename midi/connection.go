package midi

import (
	"fmt"
	"log"

	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/midi/reader"
	"gitlab.com/gomidi/rtmididrv"
)

// Interface for system to communicate with
type Interface struct {
	connected bool
	Messages  chan midi.Message
	driver    *rtmididrv.Driver
	reader    *reader.Reader
}

// NewInterface initialized
func NewInterface() *Interface {
	return &Interface{
		connected: false,
		Messages:  make(chan midi.Message),
	}
}

// Connect to the midi device
func (m *Interface) Connect(index int) {
	var err error
	if index < 0 {
		log.Println("MIDI interface not found")
		return
	}
	m.driver, err = rtmididrv.New()
	if err != nil {
		log.Println(err)
	}

	ins, err := m.driver.Ins()
	if err != nil {
		log.Println(err)
	}
	portNum := index
	in := ins[portNum]
	if err := in.Open(); err != nil {
		log.Println(err)
	}

	m.reader = reader.New(reader.NoLogger(),
		reader.Each(func(pos *reader.Position, msg midi.Message) {
			m.Messages <- msg
		}),
	)
	if err := m.reader.ListenTo(in); err != nil {
		log.Fatal("Error listening for MIDI", err)
	}
	m.connected = true
	fmt.Println("Connected to device")

}

// Disconnect from midi port
func (m *Interface) Disconnect() {
	defer m.driver.Close()
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

func (m *Interface) GetIndex(name string) int {
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
		if name == ins[i].String() {
			return i
		}
	}
	return -1
}
