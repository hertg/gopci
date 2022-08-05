package pci

import (
	"fmt"
	"os"
	"path/filepath"
)

// getDriver Get the driver by evaluating the
// 'driver' symlink inside the devices pci directory.
// Returns an empty string "" if no symlink exists.
// An error may be returned if another unexpected
// issue arises with the symlink.
func getDriver(path string) (*string, error) {
	path = filepath.Join(path, "driver")
	link, err := filepath.EvalSymlinks(path)
	if err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("unable to evaluate driver symlink '%s': %s", path, err)
	}
	if link != "" {
		link = filepath.Base(link)
	}
	return &link, nil
}
