package pci

import (
	"encoding/binary"
	"io"
)

// Resources:
// - https://wiki.osdev.org/PCI

type CommonHeader struct {
	VendorID             uint16
	DeviceID             uint16
	Command              uint16
	Status               uint16
	Revision             uint8
	ProgrammingInterface uint8
	Subclass             uint8
	Class                uint8
	CacheLine            uint8
	LatencyTimemr        uint8
	HeaderType           HeaderType
	Bist                 uint8
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

// todo: type BIST

type Config struct {
	Common CommonHeader
	Header any
}

type StandardHeader struct {
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

type PciToPciBridgeHeader struct {
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

type PciToCardbusBridgeHeader struct {
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

func ParseConfig(config io.Reader) Config {
	commonHeader := CommonHeader{}
	err := binary.Read(config, binary.LittleEndian, &commonHeader)
	if err != nil {
		panic("unable to parse common header")
	}

	if commonHeader.HeaderType.HasMultipleFunctions() {
		panic("multifunction devices are currently unsupported by this library") // fixme
	}

	var header any
	if commonHeader.HeaderType.IsGeneralDevice() {
		var d StandardHeader
		binary.Read(config, binary.LittleEndian, &d)
		header = d
	} else if commonHeader.HeaderType.IsPciToPciBridge() {
		var d PciToPciBridgeHeader
		binary.Read(config, binary.LittleEndian, &d)
		header = d
	} else if commonHeader.HeaderType.IsCardbusBridge() {
		var d PciToCardbusBridgeHeader
		binary.Read(config, binary.LittleEndian, &d)
		header = d
	} else {
		panic("unknown header type")
	}

	return Config{
		Common: commonHeader,
		Header: header,
	}
}
