package pci

import "fmt"

type Vendor struct {
	ID    uint16 `json:"id,omitempty"`
	Hex   string
	Label string `json:"label,omitempty"`
}

func ParseVendor(num uint16) *Vendor {
	vendor := Vendor{
		ID:    num,
		Hex:   fmt.Sprintf("0x%04x", num),
		Label: fmt.Sprintf("Vendor %04x", num),
	}
	if v := db.Vendors[num]; v != nil {
		vendor.Label = v.Label
	}
	return &vendor
}

func ParseSubvendor(num uint16) *Vendor {
	subvendor := Vendor{
		ID:    num,
		Hex:   fmt.Sprintf("0x%04x", num),
		Label: fmt.Sprintf("Subvendor %04x", num),
	}
	if v := db.Vendors[num]; v != nil {
		subvendor.Label = v.Label
	}
	return &subvendor
}
