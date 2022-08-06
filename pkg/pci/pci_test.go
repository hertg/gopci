package pci_test

import (
	"fmt"
	"testing"

	"github.com/hertg/gopci/pkg/pci"
)

func TestMain(t *testing.T) {

	devices, _ := pci.Scan()
	for _, device := range devices {
		fmt.Println(device.Device.Label)
	}
}
