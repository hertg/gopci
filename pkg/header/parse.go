package header

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

func Parse(reader io.Reader) IConfig {
	var buf [16]byte
	binary.Read(reader, binary.LittleEndian, &buf)
	fmt.Println(buf[:])
	headerType := HeaderType(uint8(buf[15]))
	mr := io.MultiReader(bytes.NewReader(buf[:]), reader)
	if headerType.IsGeneralDevice() {
		fmt.Println("standard")
		config := &StandardHeader{}
		binary.Read(mr, binary.LittleEndian, config)
		return config
	} else if headerType.IsPciToPciBridge() {
		fmt.Println("pci")
		config := &PciToPciBridgeHeader{}
		binary.Read(mr, binary.LittleEndian, config)
		return config
	} else if headerType.IsCardbusBridge() {
		fmt.Println("cardbus")
		config := &PciToCardbusBridgeHeader{}
		binary.Read(mr, binary.LittleEndian, config)
		return config
	} else {
		panic("unknown header type")
	}
}
