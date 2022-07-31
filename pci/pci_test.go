package pci_test

import (
	"testing"

	"github.com/hertg/gopci/pci"
)

func TestMain(t *testing.T) {
	pci.Scan()
}
