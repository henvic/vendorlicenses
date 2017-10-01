package vendorlicenses

import (
	"bytes"
	"flag"
	"io/ioutil"
	"reflect"
	"testing"
)

var update bool

func init() {
	flag.BoolVar(&update, "update", false, "update golden files")
}

func TestVendorLicensesList(t *testing.T) {
	var v = &VendorLicenses{
		Directory: ".",
	}

	var want = []string{
		"example.com/program/LICENSE",
		"x/program/license.md",
	}

	var got, err = v.List()

	if err != nil {
		t.Errorf("Expected no error, got %v instead", err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Expected %v, got %v instead", want, got)
	}
}

func TestVendorLicensesMissing(t *testing.T) {
	var v = &VendorLicenses{
		Directory: ".",
	}

	var want = []string{
		"example.com/unlicensed",
	}

	var got, err = v.Missing()

	if err != nil {
		t.Errorf("Expected no error, got %v instead", err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Expected %v, got %v instead", want, got)
	}
}

func TestVendorNoVendor(t *testing.T) {
	var v = &VendorLicenses{
		Directory: "./novendor",
	}

	var wantEmpty = []string{}

	var gotList, errList = v.List()

	if errList != nil {
		t.Errorf("Expected no error, got %v instead", wantEmpty)
	}

	if !reflect.DeepEqual(wantEmpty, gotList) {
		t.Errorf("Expected %v, got %v instead", wantEmpty, gotList)
	}

	var gotMissing, errMissing = v.Missing()

	if errMissing != nil {
		t.Errorf("Expected no error, got %v instead", wantEmpty)
	}

	if !reflect.DeepEqual(wantEmpty, gotMissing) {
		t.Errorf("Expected %v, got %v instead", wantEmpty, gotMissing)
	}
}

func TestVendorReadAll(t *testing.T) {
	var v = &VendorLicenses{
		Directory: ".",
	}

	var want = []string{
		"example.com/program/LICENSE",
		"x/program/license.md",
	}

	var got, err = v.List()

	if err != nil {
		t.Errorf("Expected no error, got %v instead", err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Expected %v, got %v instead", want, got)
	}

	var legal, errLegal = v.ReadAll()

	if errLegal != nil {
		t.Errorf("Expected no error, got %v instead", errLegal)
	}

	if update {
		if err := ioutil.WriteFile("mocks/concat", legal, 0644); err != nil {
			panic(err)
		}
	}

	b, err := ioutil.ReadFile("mocks/concat")

	if err != nil {
		panic(err)
	}

	if !bytes.Equal(legal, b) {
		t.Errorf("Expected legal text to be the same as the mock")
	}
}
