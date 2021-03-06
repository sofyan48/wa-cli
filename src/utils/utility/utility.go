package utility

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// Utility ...
type Utility struct{}

// UtilityHandler ...
func UtilityHandler() *Utility {
	return &Utility{}
}

// UtilityInterface ...
type UtilityInterface interface {
	CopyFile(src, dst string) (err error)
	ListDir(path string) ([]os.FileInfo, error)
	Rename(path, old, new string) error
}

// ListDir ..
func (util *Utility) ListDir(path string) ([]os.FileInfo, error) {
	return ioutil.ReadDir(path)
}

// Rename ...
func (util *Utility) Rename(path, old, new string) error {
	oldPath := path + "/" + old
	newPath := path + "/" + new
	return os.Rename(oldPath, newPath)
}

// CopyFile ...
func (util *Utility) CopyFile(src, dst string) (err error) {
	sfi, err := os.Stat(src)
	if err != nil {
		return
	}
	if !sfi.Mode().IsRegular() {
		return fmt.Errorf("CopyFile: non-regular source file %s (%q)", sfi.Name(), sfi.Mode().String())
	}
	dfi, err := os.Stat(dst)
	if err != nil {
		if !os.IsNotExist(err) {
			return
		}
	} else {
		if !(dfi.Mode().IsRegular()) {
			return fmt.Errorf("CopyFile: non-regular destination file %s (%q)", dfi.Name(), dfi.Mode().String())
		}
		if os.SameFile(sfi, dfi) {
			return
		}
	}
	if err = os.Link(src, dst); err == nil {
		return
	}
	err = util.copyFileContents(src, dst)
	return
}

func (util *Utility) copyFileContents(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}
