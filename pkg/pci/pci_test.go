package pci_test

import (
	"fmt"
	"testing"

	"github.com/hertg/gopci/pkg/pci"
)

func TestMain(t *testing.T) {

	classFilter := func(d *pci.Device) bool { return d.Class.Class == 0x03 }
	devices, _ := pci.Scan(classFilter)
	// devices, _ := pci.Scan(func(d *pci.Device) bool { return d.Class.Class == 0x03 })
	for _, device := range devices {
		fmt.Printf("%+v\n", device)
	}
}
