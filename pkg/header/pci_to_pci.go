package header

type PciToPciBridgeHeader struct {
	CommonHeader
	BaseAddress0                uint32
	BaseAddress1                uint32
	PrimaryBusNumber            uint8
	SecondaryBusNumber          uint8
	SubordinateBusNumber        uint8
	SecondaryLatencyTimer       uint8
	IOBase                      uint8
	IOLimit                     uint8
	SecondaryStatus             uint16
	MemoryBase                  uint16
	MemoryLimit                 uint16
	PrefetchableMemoryBase      uint16
	PrefetchableMemoryLimit     uint16
	PrefetchableBaseUpper32Bit  uint32
	PrefetchableLimitUpper32Bit uint32
	IOBaseUpper16Bit            uint16
	IOLimitUpper16Bit           uint16
	CapabilityPointer           uint8
	Reserved                    [3]byte
	ExpansionROMBaseAddress     uint32
	InterruptLine               uint8
	InterruptPin                uint8
	BridgeControl               uint16
}

func (s *PciToPciBridgeHeader) RawCommonHeader() CommonHeader {
	return s.CommonHeader
}

func (s *PciToPciBridgeHeader) DeviceID() uint16 {
	return s.CommonHeader.DeviceID
}

func (s *PciToPciBridgeHeader) VendorID() uint16 {
	return s.CommonHeader.VendorID
}

func (s *PciToPciBridgeHeader) ClassCode() uint8 {
	return s.CommonHeader.Class
}

func (s *PciToPciBridgeHeader) SubclassCode() uint8 {
	return s.CommonHeader.Subclass
}

func (s *PciToPciBridgeHeader) ProgrammingInterfaceCode() uint8 {
	return s.CommonHeader.ProgrammingInterface
}

func (s *PciToPciBridgeHeader) Revision() uint8 {
	return s.CommonHeader.Revision
}

func (s *PciToPciBridgeHeader) BIST() uint8 {
	return s.CommonHeader.BIST
}

func (s *PciToPciBridgeHeader) HeaderType() HeaderType {
	return s.CommonHeader.HeaderType
}

func (s *PciToPciBridgeHeader) LatencyTimer() uint8 {
	return s.CommonHeader.LatencyTimer
}

func (s *PciToPciBridgeHeader) CacheLineSize() uint8 {
	return s.CommonHeader.CacheLine
}
