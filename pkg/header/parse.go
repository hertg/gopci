package header

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

// Parse Expects a reader to a sysfs config file
// to be passed in and returns the parsed pci header.
// Can be of type StandardHeader, PciToPciBridgeHeader,
// or PciToCardbusBridgeHeader.
// May return an error if the provided config
// has an invalid or unknown format.
func Parse(reader io.Reader) (Config, error) {
	var buf [16]byte
	binary.Read(reader, binary.LittleEndian, &buf)
	headerType := HeaderType(uint8(buf[15]))
	mr := io.MultiReader(bytes.NewReader(buf[:]), reader)
	if headerType.IsGeneralDevice() {
		config := &StandardHeader{}
		binary.Read(mr, binary.LittleEndian, config)
		return config, nil
	} else if headerType.IsPciToPciBridge() {
		config := &PciToPciBridgeHeader{}
		binary.Read(mr, binary.LittleEndian, config)
		return config, nil
	} else if headerType.IsCardbusBridge() {
		config := &PciToCardbusBridgeHeader{}
		binary.Read(mr, binary.LittleEndian, config)
		return config, nil
	} else {
		return nil, fmt.Errorf("unknown header type: 0x%02x", headerType)
	}
}
