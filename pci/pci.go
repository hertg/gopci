package pci

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hertg/go-pciids/pkg/pciids"
)

const pciPath = "/sys/bus/pci/devices"

type Class struct {
	class    uint8
	subclass uint8
	label    string
}

type Vendor struct {
	id    uint16
	label string
}

type Device struct {
	id    uint16
	label string
}

type Subvendor struct {
	id    uint16
	label string
}

type Subdevice struct {
	id    uint16
	label string
}

type GPU struct {
	address   Address
	config    Config
	driver    string
	class     Class
	vendor    Vendor
	device    Device
	subvendor Subvendor
	subdevice Subdevice
}

var db *pciids.DB

func init() {
	filepath := "/usr/share/hwdata/pci.ids"
	file, _ := os.Open(filepath)
	scanner := bufio.NewScanner(file)
	db, _ = pciids.NewDB(scanner)
}

func Scan() ([]*GPU, error) {
	files, err := os.ReadDir(pciPath)
	if err != nil {
		panic(err)
	}

	var devices []*GPU

	for _, dir := range files {
		dev, err := ScanDevice(dir.Name())
		if err != nil {
			panic(err)
		}
		devices = append(devices, dev)
	}

	return devices, nil
}

func ScanDevice(hexAddr string) (*GPU, error) {
	path := filepath.Join(pciPath, hexAddr)
	addr := AddrFromHex(hexAddr)

	configFile, err := os.Open(filepath.Join(path, "config"))
	if err != nil {
		return nil, fmt.Errorf("unable to open config file of '%s': %s", hexAddr, err)
	}
	config := ParseConfig(configFile)

	driverPath := filepath.Join(path, "driver")
	driver, err := filepath.EvalSymlinks(driverPath)
	if err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("unable to evaluate driver symlink of '%s': %s", hexAddr, err)
	}
	if driver != "" {
		driver = filepath.Base(driver)
	}

	v := db.Vendors[pciids.VendorID(config.Common.VendorID)]
	vendor := Vendor{
		id:    uint16(v.ID),
		label: v.Label,
	}

	d := db.Devices[pciids.DeviceID(vendor.id<<16|config.Common.DeviceID)]
	device := Device{
		id:    uint16(d.ID & 0b11111111),
		label: d.Label,
	}
	fmt.Printf("%s\n%+v\n%+v\n%+v\n%+v\n\n", addr.Hex(), config, driver, vendor, device)

	return &GPU{
		address: addr,
		config:  config,
		driver:  driver,
	}, nil
}
