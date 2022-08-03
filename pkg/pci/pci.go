package pci

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hertg/go-pciids/pkg/pciids"
	"github.com/hertg/gopci/pkg/addr"
	"github.com/hertg/gopci/pkg/header"
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

type Device struct {
	Address   addr.Address
	Config    header.IConfig
	Driver    string
	Class     Class
	Vendor    Vendor
	Device    Product
	Subvendor *Vendor
	Subdevice *Product
}

var db *pciids.DB

func init() {
	filepath := "/usr/share/hwdata/pci.ids"
	file, _ := os.Open(filepath)
	scanner := bufio.NewScanner(file)
	db, _ = pciids.NewDB(scanner)
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
		fmt.Printf("%+v\n\n", dev)
		devices = append(devices, dev)
	}

	return devices, nil
}

func ScanDevice(hexAddr string) (*Device, error) {
	path := filepath.Join(pciPath, hexAddr)
	addr, _ := addr.AddrFromHex(hexAddr)

	configFile, err := os.Open(filepath.Join(path, "config"))
	if err != nil {
		return nil, fmt.Errorf("unable to open config file of '%s': %s", hexAddr, err)
	}
	defer configFile.Close()
	config := header.Parse(configFile)

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

	fmt.Println(config)

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

	// fmt.Printf("%s\t%s\n\t\t%s\n\t\t%s\n\t\t%s\n", addr.Hex(), class.Label, device.Label, vendor.label, driver)
	// fmt.Println()
	// fmt.Printf("%s\n%+v\n%+v\n%+v\n%+v\n\n", addr.Hex(), config, driver, vendor, device)

	res := &Device{
		Address: *addr,
		Config:  config,
		Driver:  driver,
		Device:  device,
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
		fmt.Println(sp)
		res.Subdevice = sp
	}

	return res, nil
}
