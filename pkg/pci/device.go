package pci

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/hertg/gopci/pkg/addr"
	"github.com/hertg/gopci/pkg/header"
)

type Device struct {
	Address   addr.Address  `json:"address,omitempty"`
	Config    header.Config `json:"config,omitempty"`
	Driver    *string       `json:"driver,omitempty"`
	Class     Class         `json:"class,omitempty"`
	Vendor    Vendor        `json:"vendor,omitempty"`
	Product   Product       `json:"product,omitempty"`
	Subvendor *Vendor       `json:"subvendor,omitempty"`
	Subdevice *Product      `json:"subdevice,omitempty"`
}

func (s *Device) SysfsPath() string {
	return filepath.Join(pciPath, s.Address.Hex())
}

func (s *Device) write(name string, b []byte) error {
	path := filepath.Join(s.SysfsPath(), name)
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0220)
	if err != nil {
		return fmt.Errorf("unable to access %s: %s", path, err)
	}
	_, err = f.Write(b)
	if err != nil {
		return fmt.Errorf("unable to write: %s", err)
	}
	return nil
}

func (s *Device) Enable() error {
	return s.write("enable", []byte("1"))
}

func (s *Device) Remove() error {
	return s.write("remove", []byte("1"))
}

func (s *Device) Rescan() error {
	return s.write("rescan", []byte("1"))
}

func (s *Device) Reset() error {
	return s.write("reset", []byte("1"))
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
