package pci

import "fmt"

type Product struct {
	ID    uint16 `json:"id,omitempty"`
	Hex   string `json:"hex,omitempty"`
	Label string `json:"label,omitempty"`
}

func ParseProduct(vendor uint16, product uint16) *Product {
	p := Product{
		ID:    product,
		Hex:   fmt.Sprintf("0x%04x", product),
		Label: fmt.Sprintf("Device %04x", product),
	}
	if label := db.FindDeviceLabel(vendor, product); label != nil {
		p.Label = *label
	}
	return &p
}

func ParseSubproduct(subvendor uint16, subproduct uint16) *Product {
	p := Product{
		ID:    subproduct,
		Hex:   fmt.Sprintf("0x%04x", subproduct),
		Label: fmt.Sprintf("Subdevice %04x", subproduct),
	}
	if label := db.FindSubsystemLabel(subvendor, subproduct); label != nil {
		p.Label = *label
	}
	return &p
}
