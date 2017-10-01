# vendorlicenses [![Build Status](http://img.shields.io/travis/henvic/vendorlicenses/master.svg?style=flat)](https://travis-ci.org/henvic/vendorlicenses) [![Coverage Status](https://coveralls.io/repos/henvic/vendorlicenses/badge.svg)](https://coveralls.io/r/henvic/vendorlicenses) [![codebeat badge](https://codebeat.co/badges/1f63b7b8-0177-4fd8-8799-9fcb7065052f)](https://codebeat.co/projects/github-com-henvic-vendorlicenses-master) [![Go Report Card](https://goreportcard.com/badge/github.com/henvic/vendorlicenses)](https://goreportcard.com/report/github.com/henvic/vendorlicenses) [![GoDoc](https://godoc.org/github.com/henvic/vendorlicenses?status.svg)](https://godoc.org/github.com/henvic/vendorlicenses)

vendorlicenses is a tool to check and concatenate licenses found on the vendor directory of Go programs.

To install it, run:

```bash
go get -u github.com/henvic/vendorlicenses/cmd/vendorlicenses
```

Usage:

```bash
$ vendorlicenses -h
Usage of vendorlicenses:
  -directory string
    	Directory to analyse (default ".")
  -list
    	List license filepaths
  -missing
    	List dependencies probably missing licenses
```

**This tool is not appropriate for vetting licenses compliance,** but to provide a concat tool for licenses similar to what is found on the legal submenus of consumer devices, such as mobile devices or cars media centers.

For vetting licensing compliance, consider using something like [FOSSA](https://fossa.io) or [FOSSology](https://www.fossology.org) instead.