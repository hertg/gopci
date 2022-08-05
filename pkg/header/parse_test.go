package header_test

import (
	"bytes"
	"testing"

	"github.com/hertg/gopci/pkg/header"
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
	expected := &header.StandardHeader{
		CommonHeader: header.CommonHeader{
			VendorID: 4098,
			DeviceID: 29631,
			Command: 1031,
			Status: 2064,
			Revision: 192,
			ProgrammingInterface: 0,
			Subclass: 0,
			Class: 3,
			CacheLine: 16,
			LatencyTimer: 0,
			HeaderType: 128,
			BIST: 0,
		},
		BaseAddress0: 12,
		BaseAddress1: 120,
		BaseAddress2: 12,
		BaseAddress3: 124,
		BaseAddress4: 57345,
		BaseAddress5: 4238344192,
		CardbusCISPointer: 0,
		SubsystemVendorID: 7854,
		SubsystemDeviceID: 26881,
		ExpansionROMBaseAddress: 4239392768,
		CapabilitiesPointer: 72,
		IRQLine: 255,
		IRQPin: 1,
		MinGnt: 0,
		MaxLat: 0,
	}
	reader := bytes.NewReader(testDevice["config"])
	config, _ := header.Parse(reader)
	c, ok := config.(*header.StandardHeader)
	assert.True(t, ok)
	assert.Equal(t, expected, c)
}
