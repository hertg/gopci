package addr_test

import (
	"testing"

	"github.com/hertg/gopci/pkg/addr"
	"github.com/stretchr/testify/assert"
)

func TestAddressParseDomain(t *testing.T) {
	a, _ := addr.AddrFromHex("0000:2f:00.0")
	assert.Equal(t, uint16(0), a.Domain())

	a, _ = addr.AddrFromHex("00ab:2f:00.0")
	assert.Equal(t, uint16(171), a.Domain())
}

func TestAddressParseBus(t *testing.T) {
	a, _ := addr.AddrFromHex("0000:2f:00.0")
	assert.Equal(t, uint8(47), a.Bus())

	a, _ = addr.AddrFromHex("00ab:02:02.3")
	assert.Equal(t, uint8(2), a.Bus())
}

func TestAddressParseDevice(t *testing.T) {
	a, _ := addr.AddrFromHex("0000:2f:00.0")
	assert.Equal(t, uint8(0), a.Device())

	a, _ = addr.AddrFromHex("0000:2f:19.0")
	assert.Equal(t, uint8(25), a.Device())

	// the max number for 5-bit is 0x1f (31)
	// any number higher than that is impossible
	// and the parser is expected to return an error
	a, err := addr.AddrFromHex("0000:2f:20.2")
	assert.Error(t, err)
}

func TestAddressParseFunction(t *testing.T) {
	a, _ := addr.AddrFromHex("0000:2f:00.0")
	assert.Equal(t, uint8(0), a.Function())

	a, _ = addr.AddrFromHex("0000:2f:00.3")
	assert.Equal(t, uint8(3), a.Function())

	// the max number for  3-bit is 0x7 (7)
	// any number higher than that is impossible
	// and the parser is expected to return an error
	a, err := addr.AddrFromHex("0000:2f:00.8")
	assert.Error(t, err)
}

func TestAddressHex(t *testing.T) {
	a, _ := addr.AddrFromHex("0002:0f:0a.2")
	assert.Equal(t, "0002:0f:0a.2", a.Hex())
}

func TestAddressHexWithoutDomain(t *testing.T) {
	a, _ := addr.AddrFromHex("2f:00.1")
	assert.Equal(t, "0000:2f:00.1", a.Hex())
}

func TestAddressParseBogusString(t *testing.T) {
	_, err := addr.AddrFromHex("abc123")
	assert.Error(t, err)

	_, err = addr.AddrFromHex("xyzx:ab:ab.0")
	assert.Error(t, err)

	// the max function num is 7, there this address is invalid
	_, err = addr.AddrFromHex("aaaa:ab:1f.8")
	assert.Error(t, err)
}
