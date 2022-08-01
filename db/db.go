package db

import (
	"bufio"
	"io"

	"github.com/hertg/gopci/internal/utils"
)

var ()

type DB struct {
	Vendors    map[VendorID]Vendor
	Devices    map[DeviceID]Device
	Subsystems map[SubsystemID]Subsystem
	Classes    map[uint8]Class
	//Subclasses            map[uint16]Subclass
	//ProgrammingInterfaces map[uint32]ProgrammingInterface
}

type VendorID uint16
type Vendor struct {
	ID      VendorID
	Label   string
	Devices map[DeviceID]*Device
	//Subsystems []*Subsystem
}

type DeviceID uint32
type Device struct {
	ID         DeviceID
	Label      string
	Subsystems map[SubsystemID]*Subsystem
}

type SubsystemID uint32
type Subsystem struct {
	ID    SubsystemID
	Label string
}

type Class struct {
	ID         uint8
	Label      string
	Subclasses map[uint8]Subclass
}

type Subclass struct {
	ID                    uint8
	Label                 string
	ProgrammingInterfaces map[uint8]ProgrammingInterface
}

type ProgrammingInterface struct {
	ID       uint8
	VendorID uint8
	DeviceID uint8
	Label    string
}

const (
	TAB       byte = 0x09
	HASH_SIGN byte = 0x23
	LETTER_C  byte = 0x43
)

func Parse(scanner *bufio.Scanner, db *DB) {

	db.Vendors = make(map[VendorID]Vendor)
	db.Devices = make(map[DeviceID]Device)
	db.Subsystems = make(map[SubsystemID]Subsystem)
	db.Classes = make(map[uint8]Class)
	//db.Subclasses = make(map[uint16]Subclass)
	//db.ProgrammingInterfaces = make(map[uint32]ProgrammingInterface)

	var vcur1 *VendorID // vendor cursor 1
	var vcur2 *DeviceID // vendor cursor 2
	var ccur1 *uint8    // class cursor 1
	var ccur2 *uint8    // class cursor 2

	for scanner.Scan() {
		line := scanner.Bytes()
		length := len(line)

		if length <= 1 {
			continue
		} else if line[0] == HASH_SIGN {
			continue
		} else if line[0] == LETTER_C {
			id := uint8(utils.ParseByteNum(line[2:4]))
			class := Class{
				ID:         id,
				Label:      string(line[6:]),
				Subclasses: make(map[uint8]Subclass, 16),
			}
			db.Classes[id] = class
			ccur1 = &class.ID
			ccur2 = nil
			vcur1 = nil
			vcur2 = nil
		} else if line[0] >= 0x30 && line[0] <= 0x39 || line[0] >= 0x61 && line[0] <= 0x66 {
			id := VendorID(utils.ParseByteNum(line[0:4]))
			vendor := Vendor{
				ID:      id,
				Label:   string(line[6:]),
				Devices: make(map[DeviceID]*Device),
			}
			db.Vendors[id] = vendor
			vcur1 = &vendor.ID
			vcur2 = nil
			ccur1 = nil
			ccur2 = nil
		} else if line[0] == TAB {
			if line[1] != TAB {
				if ccur1 != nil {
					id := uint8(utils.ParseByteNum(line[1:3]))
					subclass := Subclass{
						ID:                    id,
						Label:                 string(line[5:]),
						ProgrammingInterfaces: make(map[uint8]ProgrammingInterface),
					}
					db.Classes[*ccur1].Subclasses[id] = subclass
					ccur2 = &id
				} else if vcur1 != nil {
					id := DeviceID(utils.ParseByteNum(line[1:5]))
					device := Device{
						ID:         id,
						Label:      string(line[7:]),
						Subsystems: make(map[SubsystemID]*Subsystem),
					}
					db.Devices[id] = device
					db.Vendors[*vcur1].Devices[id] = &device
					vcur2 = &id
				} else {
					panic("got no cursor for tabbed line")
				}
			} else {
				if ccur2 != nil {
					id := uint8(utils.ParseByteNum(line[2:4]))
					progif := ProgrammingInterface{
						ID:    id,
						Label: string(line[6:]),
					}
					db.Classes[*ccur1].Subclasses[*ccur2].ProgrammingInterfaces[id] = progif
				} else if vcur2 != nil {
					subvid := uint16(utils.ParseByteNum(line[2:6]))
					subdid := uint16(utils.ParseByteNum(line[7:11]))
					id := SubsystemID(uint32(subvid)<<16 | uint32(subdid))
					subsystem := Subsystem{
						ID:    id,
						Label: string(line[13:]),
					}
					db.Subsystems[id] = subsystem
				} else {
					panic("go no cursor for double tabbed line")
				}
			}
		} else {
			continue // unexpected
		}
	}

	_ = vcur1
	_ = vcur2
	_ = ccur1
	_ = ccur2

}

func readLine(reader *bufio.Reader) {
	b, err := reader.ReadByte()
	if err != nil {
		if err != io.EOF {
			panic(err)
		}
		return
	}
	if b == 0x23 {
		// is comment
	}
	if b == 0x43 {
		// is class
	}
	if b == 0x09 {
		// is tab
	}
}
