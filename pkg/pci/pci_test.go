package pci_test

import (
	"fmt"
	"testing"

	"github.com/hertg/gopci/pkg/pci"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	// classFilter := func(d *pci.Device) bool { return d.Class.Class == 0x03 }
	devices, _ := pci.Scan()
	for _, device := range devices {
		fmt.Printf("%+v\n", device)
	}
}

func TestScanDevice(t *testing.T) {
	dev, _ := pci.ScanDeviceStr("0000:2f:00.0")
	fmt.Println(dev)

	dev, err := pci.ScanDeviceStr("0001:00:00.0")
	assert.NotNil(t, err)
}
