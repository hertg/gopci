package pci

import (
	"fmt"
	"strconv"
)

// Address Represents a PCI Address including
// domain, bus, device, and function.
type Address struct {
	number uint32
}

func (s *Address) Domain() uint16 {
	return uint16(s.number >> 16)
}

func (s *Address) Bus() uint8 {
	return uint8(s.number >> 8)
}

func (s *Address) Device() uint8 {
	return uint8(s.number>>3) & 0b00011111
}

func (s *Address) Function() uint8 {
	return uint8(s.number) & 0b00000111
}

func (s *Address) Hex() string {
	return fmt.Sprintf("%04x:%02x:%02x.%x", s.Domain(), s.Bus(), s.Device(), s.Function())
}

func addressFromString(addr string) Address {
	if len(addr) != 12 {
		panic("the pci address is expected to have a length of 12 (eg. 0000:2f:00.3)")
	}
	domain, _ := strconv.ParseUint(addr[:4], 16, 16)
	bus, _ := strconv.ParseUint(addr[5:7], 16, 8)
	device, _ := strconv.ParseUint(addr[8:10], 16, 8)
	if device > 31 {
		panic("the device can not be a number larger than 0x1f")
	}
	function, _ := strconv.ParseUint(addr[11:12], 16, 8)
	if function > 7 {
		panic("the function can not be a number larger than 0x7")
	}
	return Address{
		number: uint32((domain << 16) | (bus << 8) | (device << 3) | function),
	}
}
