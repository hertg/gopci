package pci

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddressParseDomain(t *testing.T) {
	a := addressFromString("0000:2f:00.0")
	assert.Equal(t, uint16(0), a.Domain())

	a = addressFromString("00ab:2f:00.0")
	assert.Equal(t, uint16(171), a.Domain())
}

func TestAddressParseBus(t *testing.T) {
	a := addressFromString("0000:2f:00.0")
	assert.Equal(t, uint8(47), a.Bus())

	a = addressFromString("00ab:02:02.3")
	assert.Equal(t, uint8(2), a.Bus())
}

func TestAddressParseDevice(t *testing.T) {
	a := addressFromString("0000:2f:00.0")
	assert.Equal(t, uint8(0), a.Device())

	a = addressFromString("0000:2f:19.0")
	assert.Equal(t, uint8(25), a.Device())

	// the max number for 5-bit is 0x1f (31)
	// any number higher than that is impossible
	// and the parser is expected to panic
	assert.Panics(t, func() {
		a = addressFromString("0000:2f:20.2")
	})
}

func TestAddressParseFunction(t *testing.T) {
	a := addressFromString("0000:2f:00.0")
	assert.Equal(t, uint8(0), a.Function())

	a = addressFromString("0000:2f:00.3")
	assert.Equal(t, uint8(3), a.Function())

	// the max number for  3-bit is 0x7 (7)
	// any number higher than that is impossible
	// and the parser is expected to panic
	assert.Panics(t, func() {
		a = addressFromString("0000:2f:00.8")
	})
}

func TestAddressHex(t *testing.T) {
	a := addressFromString("0002:0f:0a.2")
	assert.Equal(t, "0002:0f:0a.2", a.Hex())
}
