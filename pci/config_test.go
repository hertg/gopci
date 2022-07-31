package pci_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/hertg/gopci/pci"
	"github.com/stretchr/testify/assert"
)

// note: hexdump -ve '1/1 "0x%.2x,"' config

var testDevice = map[string][]byte{
	"class": []byte("0x030000"),
	"config": {
		0x02, 0x10, 0xbf, 0x73, 0x07, 0x04, 0x10, 0x08,
		0xc0, 0x00, 0x00, 0x03, 0x10, 0x00, 0x80, 0x00,
		0x0c, 0x00, 0x00, 0x00, 0x78, 0x00, 0x00, 0x00,
		0x0c, 0x00, 0x00, 0x00, 0x7c, 0x00, 0x00, 0x00,
		0x01, 0xe0, 0x00, 0x00, 0x00, 0x00, 0xa0, 0xfc,
		0x00, 0x00, 0x00, 0x00, 0xae, 0x1e, 0x01, 0x69,
		0x00, 0x00, 0xb0, 0xfc, 0x48, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0xff, 0x01, 0x00, 0x00,
	},
	"device":           []byte("0x73bf"),
	"enable":           []byte("1"),
	"irq":              []byte("197"),
	"local_cpus":       []byte("ffffffff"),
	"local_cpulist":    []byte("0-31"),
	"resource":         {},
	"revision":         []byte("0xc0"),
	"rom":              {},
	"subsystem_device": []byte("0x6901"),
	"subsystem_vendor": []byte("0x1eae"),
	"vendor":           []byte("0x1002"),
}

func TestConfigParse(t *testing.T) {
	reader := bytes.NewReader(testDevice["config"])
	config := pci.ParseConfig(reader)

	assert.Equal(t, uint8(0x03), config.Common.Class)
	assert.Equal(t, uint8(0x00), config.Common.Subclass)
	assert.Equal(t, uint16(0x1002), config.Common.VendorID)
	assert.Equal(t, uint16(0x73bf), config.Common.DeviceID)
	assert.Equal(t, uint8(0xc0), config.Common.Revision)

	fmt.Printf("%+v\n", config)
}
