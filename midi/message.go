package midi

// Message to/from midi device
type Message struct {
	Value    string
	Channel  uint8
	Key      uint8
	Velocity uint8
}
