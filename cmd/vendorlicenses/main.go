package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/henvic/vendorlicenses"
)

var (
	directory string
	list      bool
	missing   bool
)

func init() {
	flag.BoolVar(&list, "list", false, "List license filepaths")
	flag.StringVar(&directory, "directory", ".", "Directory to analyze")
	flag.BoolVar(&missing, "missing", false, "List dependencies probably missing licenses")
}

func main() {
	flag.Parse()

	var v = &vendorlicenses.VendorLicenses{
		Directory: directory,
	}

	var err error

	switch {
	case list:
		err = listV(v)
	case missing:
		err = missingV(v)
	default:
		err = readAllV(v)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func listV(v *vendorlicenses.VendorLicenses) error {
	paths, err := v.List()

	if err != nil {
		return err
	}

	for _, p := range paths {
		fmt.Println(p)
	}

	return nil
}

func missingV(v *vendorlicenses.VendorLicenses) error {
	missingLicenses, err := v.Missing()

	if err != nil {
		return err
	}

	for _, p := range missingLicenses {
		fmt.Println(p)
	}

	return nil
}

func readAllV(v *vendorlicenses.VendorLicenses) error {
	licenses, err := v.ReadAll()

	if err != nil {
		return err
	}

	fmt.Print(string(licenses))
	return nil
}
