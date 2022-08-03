package header

// Resources:
// - https://wiki.osdev.org/PCI

type IConfig interface {
	RawCommonHeader() CommonHeader
	DeviceID() uint16
	VendorID() uint16
	ClassCode() uint8
	SubclassCode() uint8
	ProgrammingInterfaceCode() uint8
	Revision() uint8
	BIST() uint8
	HeaderType() HeaderType
	LatencyTimer() uint8
	CacheLineSize() uint8
}
