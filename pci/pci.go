package pci

import (
	"fmt"
	"os"
	"path/filepath"
)

const pciPath = "/sys/bus/pci/devices"

type Device struct {
	address Address
	config  Config
	driver  string
}

func Scan() ([]*Device, error) {
	files, err := os.ReadDir(pciPath)
	if err != nil {
		panic(err)
	}

	var devices []*Device

	for _, dir := range files {
		dev, err := ScanDevice(dir.Name())
		if err != nil {
			panic(err)
		}
		devices = append(devices, dev)
	}

	return devices, nil
}

func ScanDevice(hexAddr string) (*Device, error) {
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

	fmt.Printf("%s\n%+v\n%+v\n\n", addr.Hex(), config, driver)

	return &Device{
		address: addr,
		config:  config,
		driver:  driver,
	}, nil
}
