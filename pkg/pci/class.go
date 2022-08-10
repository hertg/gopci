package pci

import "fmt"

type Class struct {
	ID    uint32 `json:"id,omitempty"`
	Hex   string `json:"hex,omitempty"`
	Label string `json:"label,omitempty"`
}

func ParseClass(class uint8, subclass uint8, progif uint8) *Class {
	combined := uint32(class)<<16 | uint32(subclass)<<8 | uint32(progif)
	c := Class{
		ID:    combined,
		Hex:   fmt.Sprintf("0x%06x", combined),
		Label: fmt.Sprintf("Class %06x", combined),
	}
	if label := db.FindProgifLabel(class, subclass, progif); label != nil {
		c.Label = *label
	}
	return &c
}
