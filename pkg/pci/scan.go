package pci

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hertg/gopci/pkg/addr"
	"github.com/hertg/gopci/pkg/header"
)

var (
	ErrDeviceNotFound error = errors.New("pci device not found")
	ErrNoConfig       error = errors.New("no pci 'config' file found")
)

// ScanDevice Scan the device at a specific PCI address
// and return the parsed information. May return
// ErrDeviceNotFound if no device can be found at the provided
// PCI address, or ErrNoConfig if the device's PCI path
// contains no 'config' file.
func ScanDevice(addr *addr.Address) (*Device, error) {
	path := filepath.Join(pciPath, addr.Hex())
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return nil, ErrDeviceNotFound
	}

	configFile, err := os.Open(filepath.Join(path, "config"))
	if err != nil {
		return nil, ErrNoConfig
	}
	defer configFile.Close()
	config, err := header.Parse(configFile)
	if err != nil {
		return nil, err
	}

	driver, err := getDriver(path)
	if err != nil {
		return nil, err
	}

	class := Class{
		Class:    config.ClassCode(),
		Subclass: config.SubclassCode(),
		Label:    fmt.Sprintf("Class %x", config.ClassCode()),
	}
	cl := db.FindSubclass(uint16(config.ClassCode())<<8 | uint16(config.SubclassCode()))
	if cl != nil {
		class.Label = *cl
	}

	v := db.Vendors[config.VendorID()]
	vendor := Vendor{
		ID:    config.VendorID(),
		label: fmt.Sprintf("Vendor %2.x", config.VendorID()),
	}
	if v != nil {
		vendor.label = v.Label
	}

	d := db.Devices[uint32(vendor.ID)<<16|uint32(config.DeviceID())]
	device := Product{
		ID:    config.DeviceID(),
		Label: fmt.Sprintf("Device %2.x", config.DeviceID()),
	}
	if d != nil {
		device.Label = d.Label
	}

	res := &Device{
		Address: *addr,
		Config:  config,
		Driver:  *driver,
		Product: device,
		Vendor:  vendor,
		Class:   class,
	}

	if c, ok := config.(*header.StandardHeader); ok {
		sv := &Vendor{
			ID:    c.SubsystemVendorID,
			label: fmt.Sprintf("Subvendor %04x", c.SubsystemVendorID),
		}
		res.Subvendor = sv

		sp := &Product{
			ID:    c.SubsystemDeviceID,
			Label: fmt.Sprintf("Subdevice %04x", c.SubsystemDeviceID),
		}
		res.Subdevice = sp
	}

	return res, nil
}

func ScanDeviceStr(hexAddr string) (*Device, error) {
	addr, err := addr.AddrFromHex(hexAddr)
	if err != nil {
		return nil, err
	}
	return ScanDevice(addr)
}
