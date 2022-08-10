package header

// Resources:
// - https://wiki.osdev.org/PCI

type Config interface {
	// RawCommonHeader Retrieve the raw common header fields
	// that are present in all header types
	RawCommonHeader() CommonHeader

	// DeviceID Identifies the particular device.
	// Where valid IDs are allocated by the vendor.
	DeviceID() uint16

	// VendorID Identifies the manufacturer of the device.
	// Where valid IDs are allocated by PCI-SIG to
	// ensure uniqueness and 0xFFFF is an invalid value
	// that will be returned on read accesses to
	// Configuration Space registers of non-existent devices.
	VendorID() uint16

	// Status A register used to record status
	// information for PCI bus related events.
	Status() uint16

	// Command Provides control over a device's ability
	// to generate and respond to PCI cycles.
	Command() uint16

	// ClassCode A read-only register that specifies
	// the type of function the device performs.
	ClassCode() uint8

	// SubclassCode A read-only register that specifies
	// the specific function the device performs.
	SubclassCode() uint8

	// ProgrammingInterfaceCode A read-only register that
	// specifies a register-level programming interface
	// the device has, if it has any at all.
	ProgrammingInterfaceCode() uint8

	// Revision Specifies a revision identifier for a particular device.
	// Where valid IDs are allocated by the vendor.
	Revision() uint8

	// BIST Represents that status and allows control
	// of a devices BIST (built-in self test).
	BIST() uint8

	// HeaderType Identifies the layout of the rest of
	// the header beginning at byte 0x10 of the header
	// and also specifies whether or not the device
	// has multiple functions.
	HeaderType() HeaderType

	// LatencyTimer Specifies the latency timer in units of PCI bus clocks.
	LatencyTimer() uint8

	// CacheLineSize Specifies the system cache line size in 32-bit units.
	CacheLineSize() uint8
}
