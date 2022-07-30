package pci

import (
	"fmt"
	"os"
	"path/filepath"
)

const pciPath = "/sys/bus/pci/devices"

type Device struct {
	address Address
}

func Scan() {

	//	var devices []Device

	files, err := os.ReadDir(pciPath)
	if err != nil {
		panic(err)
	}

	for _, dir := range files {
		addr := addressFromString(dir.Name())
		link, err := filepath.EvalSymlinks(filepath.Join(pciPath, dir.Name()))
		if err != nil {
			panic(err)
		}
		fmt.Println(addr, addr.Domain(), addr.Bus(), addr.Device(), addr.Function())
		fmt.Println(link)
	}
}
