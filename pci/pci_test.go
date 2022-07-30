package pci_test

import (
	"testing"

	"github.com/hertg/go-readpci/pci"
)

func TestMain(t *testing.T) {
	pci.Scan()
}
