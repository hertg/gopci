package pci

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/hertg/gopci/pkg/addr"
	"github.com/hertg/gopci/pkg/header"
)

type Device struct {
	Address   addr.Address
	Config    header.IConfig
	Driver    string
	Class     Class
	Vendor    Vendor
	Product   Product
	Subvendor *Vendor
	Subdevice *Product
}

func (s *Device) SysfsPath() string {
	return filepath.Join(pciPath, s.Address.Hex())
}

func (s *Device) write(name string, b []byte) error {
	path := filepath.Join(s.SysfsPath(), name)
	f, err := os.OpenFile(path, os.O_TRUNC, 644)
	if err != nil {
		return fmt.Errorf("unable to access %s: %s", path, err)
	}
	_, err = f.Write(b)
	if err != nil {
		return fmt.Errorf("unable to enable device: %s", err)
	}
	return nil
}

func (s *Device) Enable() error {
	return s.write("enable", []byte{1})
}

func (s *Device) Remove() error {
	return s.write("remove", []byte{1})
}

func (s *Device) Rescan() error {
	return s.write("rescan", []byte{1})
}

func (s *Device) Reset() error {
	return s.write("reset", []byte{1})
}

// todo: current_link_speed
// todo: current_link_width
// todo: max_link_speed
// todo: max_link_width
// todo: numa_node
// todo: local_cpulist
// todo: local_cpus
// todo: hwmon
// todo: serial_number (?)
// todo: vbios_version
