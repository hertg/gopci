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
	Class    uint8
	Subclass uint8
	Label    string
}

type Vendor struct {
	ID    uint16
	label string
}

type Device struct {
	ID    uint16
	Label string
}

type Subvendor struct {
	ID    uint16
	Label string
}

type Subdevice struct {
	ID    uint16
	Label string
}

type GPU struct {
	Address   Address
	Config    Config
	Driver    string
	Class     Class
	Vendor    Vendor
	Device    Device
	Subvendor Subvendor
	Subdevice Subdevice
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
	addr, _ := AddrFromHex(hexAddr)

	configFile, err := os.Open(filepath.Join(path, "config"))
	if err != nil {
		return nil, fmt.Errorf("unable to open config file of '%s': %s", hexAddr, err)
	}
	defer configFile.Close()
	config := ParseConfig(configFile)

	driverPath := filepath.Join(path, "driver")
	driver, err := filepath.EvalSymlinks(driverPath)
	if err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("unable to evaluate driver symlink of '%s': %s", hexAddr, err)
	}
	if driver != "" {
		driver = filepath.Base(driver)
	} else {
		driver = "-"
	}

	class := Class{
		Class:    config.Common.Class,
		Subclass: config.Common.Subclass,
		Label:    fmt.Sprintf("Class %x", config.Common.Class),
	}
	cl := db.FindSubclass(uint16(config.Common.Class)<<8 | uint16(config.Common.Subclass))
	if cl != nil {
		class.Label = *cl
	}

	v := db.Vendors[config.Common.VendorID]
	vendor := Vendor{
		ID:    config.Common.VendorID,
		label: fmt.Sprintf("Vendor %2.x", config.Common.VendorID),
	}
	if v != nil {
		vendor.label = v.Label
	}

	d := db.Devices[uint32(vendor.ID)<<16|uint32(config.Common.DeviceID)]
	device := Device{
		ID:    config.Common.DeviceID,
		Label: fmt.Sprintf("Device %2.x", config.Common.DeviceID),
	}
	if d != nil {
		device.Label = d.Label
	}

	// fmt.Printf("%s\t%s\n\t\t%s\n\t\t%s\n\t\t%s\n", addr.Hex(), class.Label, device.Label, vendor.label, driver)
	// fmt.Println()
	// fmt.Printf("%s\n%+v\n%+v\n%+v\n%+v\n\n", addr.Hex(), config, driver, vendor, device)

	return &GPU{
		Address: *addr,
		Config:  config,
		Driver:  driver,
	}, nil
}
