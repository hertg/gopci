package pci_test

import (
	"testing"

	"github.com/hertg/gopci/pci"
	"github.com/jaypipes/ghw"
)

var devices any

func BenchmarkGoPci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		devices, _ = pci.Scan()
	}
}

func BenchmarkGhw(b *testing.B) {
	for i := 0; i < b.N; i++ {
		devices, _ = ghw.PCI()
	}
}
