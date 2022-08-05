package header

type PciToCardbusBridgeHeader struct {
	CommonHeader
	CardbusSocketBaseAddress uint32
	OffsetOfCapabilitiesList uint8
	Reserved                 byte
	SecondaryStatus          uint16
	PciBusNumber             uint8
	CardbusBusNumber         uint8
	SubordinateBusNumber     uint8
	CardbusLatencyTimer      uint8
	MemoryBaseAddress0       uint32
	MemoryLimit0             uint32
	MemoryBaseAddress1       uint32
	MemoryLimit1             uint32
	IOBaseAddress0           uint32
	IOLimit0                 uint32
	IOBaseAddress1           uint32
	IOLimit1                 uint32
	InterrupLine             uint8
	InterruptPin             uint8
	BridgeControl            uint16
	SubsystemDeviceID        uint16
	SubsystemVendorID        uint16
	LegacyModeBaseAddress    uint32
}

func (s *PciToCardbusBridgeHeader) RawCommonHeader() CommonHeader {
	return s.CommonHeader
}

func (s *PciToCardbusBridgeHeader) DeviceID() uint16 {
	return s.CommonHeader.DeviceID
}

func (s *PciToCardbusBridgeHeader) VendorID() uint16 {
	return s.CommonHeader.VendorID
}

func (s *PciToCardbusBridgeHeader) Status() uint16 {
	return s.CommonHeader.Status
}

func (s *PciToCardbusBridgeHeader) Command() uint16 {
	return s.CommonHeader.Command
}

func (s *PciToCardbusBridgeHeader) ClassCode() uint8 {
	return s.CommonHeader.Class
}

func (s *PciToCardbusBridgeHeader) SubclassCode() uint8 {
	return s.CommonHeader.Subclass
}

func (s *PciToCardbusBridgeHeader) ProgrammingInterfaceCode() uint8 {
	return s.CommonHeader.ProgrammingInterface
}

func (s *PciToCardbusBridgeHeader) Revision() uint8 {
	return s.CommonHeader.Revision
}

func (s *PciToCardbusBridgeHeader) BIST() uint8 {
	return s.CommonHeader.BIST
}

func (s *PciToCardbusBridgeHeader) HeaderType() HeaderType {
	return s.CommonHeader.HeaderType
}

func (s *PciToCardbusBridgeHeader) LatencyTimer() uint8 {
	return s.CommonHeader.LatencyTimer
}

func (s *PciToCardbusBridgeHeader) CacheLineSize() uint8 {
	return s.CommonHeader.CacheLine
}
