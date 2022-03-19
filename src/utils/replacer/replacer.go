package replacer

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Replacer struct{}

// ReplacerHandler ...
func ReplacerHandler() *Replacer {
	return &Replacer{}
}

type ReplacerInterface interface {
	Replace(project, old, new string) error
}

// Replace ...
func (r *Replacer) Replace(project, old, new string) error {
	setup := func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !!fi.IsDir() {
			return nil
		}
		matched, err := filepath.Match("*", fi.Name())
		if err != nil {
			return err
		}

		if matched {
			read, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			orgName := strings.Replace(string(read), old, new, -1)
			err = ioutil.WriteFile(path, []byte(orgName), 0)
			if err != nil {
				return err
			}
		}

		return nil
	}
	return filepath.Walk("./"+project, setup)
}
