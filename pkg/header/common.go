package header

type CommonHeader struct {
	VendorID             uint16 // required
	DeviceID             uint16 // required
	Command              uint16 // required
	Status               uint16 // required
	Revision             uint8  // required
	ProgrammingInterface uint8
	Subclass             uint8
	Class                uint8 // required
	CacheLine            uint8
	LatencyTimer         uint8
	HeaderType           HeaderType // required
	BIST                 uint8
}

type HeaderType uint8

func (s HeaderType) HasMultipleFunctions() bool {
	// returns whether the 7th bit is set
	return (s & (1 << (7 - 1))) != 0
}

func (s HeaderType) IsGeneralDevice() bool {
	return s&0b01111111 == 0x0 // mask the multifunction bit
}

func (s HeaderType) IsPciToPciBridge() bool {
	return s&0b01111111 == 0x1 // mask the multifunction bit
}

func (s HeaderType) IsCardbusBridge() bool {
	return s&0b01111111 == 0x2 // mask the multifunction bit
}
