package vendorlicenses

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// VendorLicenses parses the vendor directory to find licenses
type VendorLicenses struct {
	Directory string
	vendors   string
	list      []string
	missing   []string
	listed    bool
	buffer    bytes.Buffer
}

// List and load paths to license files on the vendor directory
func (v *VendorLicenses) List() (list []string, err error) {
	v.listed = true
	v.list = []string{}
	v.vendors, err = filepath.Abs(filepath.Join(v.Directory, "vendor"))

	if err != nil {
		return nil, err
	}

	if _, err = os.Stat(v.vendors); os.IsNotExist(err) {
		return []string{}, nil
	}

	err = filepath.Walk(v.vendors, v.walkFn)
	return v.list, err
}

// ReadAll reads all licenses on the vendor directory
func (v *VendorLicenses) ReadAll() (legal []byte, err error) {
	if err := v.maybeList(); err != nil {
		return nil, err
	}

	for n, p := range v.list {
		if err = v.read(p); err != nil {
			return v.buffer.Bytes(), err
		}

		if n < len(v.list)-1 {
			v.buffer.WriteString("\n\n")
		}
	}

	return v.buffer.Bytes(), nil
}

func (v *VendorLicenses) read(path string) error {
	var dirPath = filepath.Dir(path)

	v.buffer.WriteString(fmt.Sprintf(
		"%s %s\n",
		filepath.Base(dirPath),
		dirPath))

	text, err := ioutil.ReadFile(
		filepath.Join(v.vendors, path),
	)

	if err != nil {
		return err
	}

	v.buffer.Write(text)
	return nil
}

func (v *VendorLicenses) walkFn(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if info.IsDir() {
		return nil
	}

	// files ending in .go are most likely not licenses at all
	if strings.HasPrefix(strings.ToLower(info.Name()), "license") &&
		!strings.HasSuffix(info.Name(), ".go") {
		rel, _ := filepath.Rel(v.vendors, path)
		v.list = append(v.list, rel)
	}

	return nil
}

// Missing licenses on a given vendor directory
func (v *VendorLicenses) Missing() (missing []string, err error) {
	if err = v.maybeList(); err != nil {
		return nil, err
	}

	if _, err = os.Stat(v.vendors); os.IsNotExist(err) {
		return []string{}, nil
	}

	err = filepath.Walk(v.vendors, v.walkMissingFn)
	return v.missing, err
}

func (v *VendorLicenses) walkMissingFn(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if !info.IsDir() {
		return nil
	}

	// check if size of dir is empty.......
	for _, licensePath := range v.list {
		rel, err := filepath.Rel(
			filepath.Join(v.vendors, filepath.Dir(licensePath)), path)

		if err != nil {
			return err
		}

		if !strings.Contains(rel, "..") {
			return filepath.SkipDir
		}
	}

	has, err := checkIfDirHasFilesOnFirstLevel(path)

	if err != nil {
		return err
	}

	if has {
		rel, _ := filepath.Rel(v.vendors, path)
		v.missing = append(v.missing, rel)
	}

	return nil
}

func (v *VendorLicenses) maybeList() error {
	if v.listed {
		return nil
	}

	_, err := v.List()
	return err
}

// check if directory has files on first level
// it ignores files starting with . (dot)
func checkIfDirHasFilesOnFirstLevel(dir string) (bool, error) {
	fis, err := ioutil.ReadDir(dir)

	if err != nil {
		return false, err
	}

	for _, fi := range fis {
		if fi.IsDir() || fi.Size() == 0 {
			continue
		}

		if !strings.HasPrefix(fi.Name(), ".") {
			return true, nil
		}
	}

	return false, nil
}
