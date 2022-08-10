package pci

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/hertg/gopci/pkg/addr"
	"github.com/hertg/gopci/pkg/header"
)

var (
	ErrDeviceNotFound error = errors.New("pci device not found")
	ErrNoConfig       error = errors.New("no pci 'config' file found")
	ErrInvalidConfig  error = errors.New("the pci 'config' file has an invalid format")
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
		return nil, ErrInvalidConfig
	}

	driver, err := getDriver(path)
	if err != nil {
		return nil, err
	}
	class := ParseClass(config.ClassCode(), config.SubclassCode(), config.ProgrammingInterfaceCode())
	vendor := ParseVendor(config.VendorID())
	device := ParseProduct(config.VendorID(), config.DeviceID())
	ret := &Device{
		Address: *addr,
		Config:  config,
		Driver:  driver,
		Product: *device,
		Vendor:  *vendor,
		Class:   *class,
	}
	if c, ok := config.(*header.StandardHeader); ok {
		ret.Subvendor = ParseSubvendor(c.SubsystemVendorID)
		ret.Subdevice = ParseSubproduct(c.SubsystemVendorID, c.SubsystemDeviceID)
	}
	return ret, nil
}

func ScanDeviceStr(hexAddr string) (*Device, error) {
	addr, err := addr.AddrFromHex(hexAddr)
	if err != nil {
		return nil, err
	}
	return ScanDevice(addr)
}
