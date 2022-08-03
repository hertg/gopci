package header

import (
	"bytes"
	"encoding/binary"
	"io"
)

func Parse(reader io.Reader) IConfig {
	var buf [16]byte
	binary.Read(reader, binary.LittleEndian, &buf)
	headerType := HeaderType(uint8(buf[14]))
	mr := io.MultiReader(bytes.NewReader(buf[:]), reader)
	if headerType.IsGeneralDevice() {
		config := &StandardHeader{}
		binary.Read(mr, binary.LittleEndian, config)
		return config
	} else if headerType.IsPciToPciBridge() {
		config := &PciToPciBridgeHeader{}
		binary.Read(mr, binary.LittleEndian, config)
		return config
	} else if headerType.IsCardbusBridge() {
		config := &PciToCardbusBridgeHeader{}
		binary.Read(mr, binary.LittleEndian, config)
		return config
	} else {
		panic("unknown header type")
	}
}
