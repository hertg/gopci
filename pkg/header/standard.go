package header

type StandardHeader struct {
	CommonHeader
	BaseAddress0            uint32
	BaseAddress1            uint32
	BaseAddress2            uint32
	BaseAddress3            uint32
	BaseAddress4            uint32
	BaseAddress5            uint32
	CardbusCISPointer       uint32
	SubsystemVendorID       uint16
	SubsystemDeviceID       uint16
	ExpansionROMBaseAddress uint32
	CapabilitiesPointer     uint8
	Reserved                [7]byte
	IRQLine                 uint8
	IRQPin                  uint8
	MinGnt                  uint8
	MaxLat                  uint8
}

func (s *StandardHeader) RawCommonHeader() CommonHeader {
	return s.CommonHeader
}

func (s *StandardHeader) DeviceID() uint16 {
	return s.CommonHeader.DeviceID
}

func (s *StandardHeader) VendorID() uint16 {
	return s.CommonHeader.VendorID
}

func (s *StandardHeader) ClassCode() uint8 {
	return s.CommonHeader.Class
}

func (s *StandardHeader) SubclassCode() uint8 {
	return s.CommonHeader.Subclass
}

func (s *StandardHeader) ProgrammingInterfaceCode() uint8 {
	return s.CommonHeader.ProgrammingInterface
}

func (s *StandardHeader) Revision() uint8 {
	return s.CommonHeader.Revision
}

func (s *StandardHeader) BIST() uint8 {
	return s.CommonHeader.BIST
}

func (s *StandardHeader) HeaderType() HeaderType {
	return s.CommonHeader.HeaderType
}

func (s *StandardHeader) LatencyTimer() uint8 {
	return s.CommonHeader.LatencyTimer
}

func (s *StandardHeader) CacheLineSize() uint8 {
	return s.CommonHeader.CacheLine
}
