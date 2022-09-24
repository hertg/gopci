package addr

import (
	"fmt"
	"strconv"
)

// Address Represents a PCI Address including
// domain, bus, device, and function.
type Address struct {
	// Number The 32bit/48bit PCI hardware address.
	// This contains the domain (16/32 bits), bus (8 bits),
	// device (5 bits), and function (3 bits).
	Number uint64
}

// Domain Get the domain as uint32
func (s *Address) Domain() uint32 {
	return uint32(s.Number >> 16)
}

// DomainHex Get the domain as hexadecimal string.
// Example: '0000', '10000', or 'deadbeef'
func (s *Address) DomainHex() string {
	if s.Domain() > 65535 {
		// if domain is larger than 0xffff, don't cap output at 4 hex digits
		return fmt.Sprintf("%x", s.Domain())
	}
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
	return fmt.Sprintf("%x", s.Function())
}

// Hex Get the PCI Address in full human readable form.
// This includes the domain, bus, device, and function.
// Example: 0000:2f:00.1, or 10000:2f:00.1
func (s *Address) Hex() string {
	if s.Domain() > 65535 {
		return fmt.Sprintf("%x:%02x:%02x.%x", s.Domain(), s.Bus(), s.Device(), s.Function())
	}
	return fmt.Sprintf("%04x:%02x:%02x.%x", s.Domain(), s.Bus(), s.Device(), s.Function())
}

// AddrFromHex Parse a human readable PCI address
// into an Address struct. Omitting the domain
// defaults to '0000'.
//
// Expected formats: '10000:2f:00.1', '0000:2f:00.1', or '2f:00.1'
// -> The domain is expected to be omitted OR 4-8 chars long
func AddrFromHex(addr string) (*Address, error) {
	if len(addr) == 7 {
		addr = "0000:" + addr
	}
	domainLength := len(addr) - 8
	if domainLength < 4 || domainLength > 8 {
		return nil, fmt.Errorf("unable to parse '%s' as pci address", addr)
	}
	domain, err := strconv.ParseUint(addr[:domainLength], 16, 32)
	if err != nil {
		return nil, fmt.Errorf("unable to parse pci domain address")
	}
	addr = addr[domainLength+1:]
	bus, err := strconv.ParseUint(addr[:2], 16, 8)
	if err != nil {
		return nil, fmt.Errorf("unable to parse pci bus address '%s'", addr[:2])
	}
	addr = addr[3:]
	device, err := strconv.ParseUint(addr[:2], 16, 8)
	if err != nil {
		return nil, fmt.Errorf("unable to parse pci device address '%s'", addr[:2])
	}
	if device > 31 {
		return nil, fmt.Errorf("the device can not be a number larger than 0x1f")
	}
	addr = addr[3:]
	function, err := strconv.ParseUint(addr[:1], 16, 8)
	if err != nil {
		return nil, fmt.Errorf("unable to parse pci bus function '%s'", addr[3:])
	}
	if function > 7 {
		return nil, fmt.Errorf("the function can not be a number larger than 0x7")
	}
	return &Address{
		Number: uint64((domain << 16) | (bus << 8) | (device << 3) | function),
	}, nil
}
