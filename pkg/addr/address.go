package addr

import (
	"fmt"
	"strconv"
)

// Address Represents a PCI Address including
// domain, bus, device, and function.
type Address struct {
	// Number The 32-bit PCI hardware address.
	// This contains the domain (16 bits), bus (8 bits),
	// device (5 bits), and function (3 bits).
	Number uint32
}

// Domain Get the domain as uint16
func (s *Address) Domain() uint16 {
	return uint16(s.Number >> 16)
}

// DomainHex Get the domain as hexadecimal string.
// Example: '0000'
func (s *Address) DomainHex() string {
	return fmt.Sprintf("%04x", s.Domain())
}

// Bus Get the bus as uint8
func (s *Address) Bus() uint8 {
	return uint8(s.Number >> 8)
}

// BusHex Get the bus as hexadecimal string.
// Example: '2f'
func (s *Address) BusHex() string {
	return fmt.Sprintf("%02x", s.Bus())
}

// Device Get the device as uint8.
// The number returned is effectively 5-bit,
// and is therefore always between 0 to 31.
func (s *Address) Device() uint8 {
	return uint8(s.Number>>3) & 0b00011111
}

// DeviceHex Get the device as hexadecimal string.
// Example: '00'
func (s *Address) DeviceHex() string {
	return fmt.Sprintf("%02x", s.Device())
}

// Function Get the function as uint8.
// The number returned is effectively 3-bit,
// and is therefore always between 0 to 7.
func (s *Address) Function() uint8 {
	return uint8(s.Number) & 0b00000111
}

// FunctionHex Get the function as hexadecimal string.
// Example: '1'
func (s *Address) FunctionHex() string {
	return fmt.Sprintf("%x", s.Device())
}

// Hex Get the PCI Address in full human readable form.
// This includes the domain, bus, device, and function.
// Example: 0000:2f:00.1
func (s *Address) Hex() string {
	return fmt.Sprintf("%04x:%02x:%02x.%x", s.Domain(), s.Bus(), s.Device(), s.Function())
}

// AddrFromHex Parse a human readable PCI address
// into an Address struct. Omitting the domain
// defaults to '0000'.
// Expected formats: '0000:2f:00.1' or '2f:00.1'
func AddrFromHex(addr string) (*Address, error) {
	if len(addr) == 7 {
		addr = "0000:" + addr
	}
	if len(addr) != 12 {
		return nil, fmt.Errorf("unable to parse '%s' as pci address", addr)
	}
	domain, err := strconv.ParseUint(addr[:4], 16, 16)
	if err != nil {
		return nil, fmt.Errorf("unable to parse pci domain address")
	}
	bus, err := strconv.ParseUint(addr[5:7], 16, 8)
	if err != nil {
		return nil, fmt.Errorf("unable to parse pci bus address")
	}
	device, err := strconv.ParseUint(addr[8:10], 16, 8)
	if err != nil {
		return nil, fmt.Errorf("unable to parse pci device address")
	}
	if device > 31 {
		return nil, fmt.Errorf("the device can not be a number larger than 0x1f")
	}
	function, _ := strconv.ParseUint(addr[11:12], 16, 8)
	if function > 7 {
		return nil, fmt.Errorf("the function can not be a number larger than 0x7")
	}
	return &Address{
		Number: uint32((domain << 16) | (bus << 8) | (device << 3) | function),
	}, nil
}
