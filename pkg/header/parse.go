package header

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

func Parse(reader io.Reader) (IConfig, error) {
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
		return nil, fmt.Errorf("unknown header type: 0x%02X", headerType)
	}
}
