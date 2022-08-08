package pci

import (
	"bufio"
	"fmt"
	"os"

	"github.com/hertg/go-pciids/pkg/pciids"
)

const pciPath = "/sys/bus/pci/devices"

type Class struct {
	Class    uint8
	Subclass uint8
	Label    string
}

type Vendor struct {
	ID    uint16
	label string
}

type Product struct {
	ID    uint16
	Label string
}

var db *pciids.DB

func init() {
	filepath := "/usr/share/hwdata/pci.ids"
	file, _ := os.Open(filepath)
	scanner := bufio.NewScanner(file)
	db, _ = pciids.NewDB(scanner)
}

func Scan(filters ...func(*Device) bool) ([]*Device, error) {
	files, err := os.ReadDir(pciPath)
	if err != nil {
		return nil, fmt.Errorf("got error while reading '%s': %s", pciPath, err)
	}
	var devices []*Device
	for _, dir := range files {
		dev, err := ScanDeviceStr(dir.Name())
		if err != nil {
			return nil, fmt.Errorf("got error while scanning device '%s': %s", dir.Name(), err)
		}
		skip := false
		for _, filter := range filters {
			if !filter(dev) {
				skip = true
				continue
			}
		}
		if !skip {
			devices = append(devices, dev)
		}
	}
	return devices, nil
}
