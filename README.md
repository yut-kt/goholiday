# goholiday

[![v1.0.4](https://img.shields.io/github/v/release/yut-kt/goholiday?logoColor=ff69b4&style=social)]()
[![Test](https://github.com/yut-kt/goholiday/actions/workflows/default_branch_test.yaml/badge.svg)](https://github.com/yut-kt/goholiday/actions/workflows/default_branch_test.yaml)
[![coverage](https://img.shields.io/badge/coverage-100%25-green.svg)]()
[![Go Report Card](https://goreportcard.com/badge/github.com/yut-kt/goholiday)](https://goreportcard.com/report/github.com/yut-kt/goholiday)  
[![GoDoc](https://godoc.org/github.com/yut-kt/goholiday?status.svg)](https://godoc.org/github.com/yut-kt/goholiday)
[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](LICENSE)

**Calculating (holiday, business) days support for golang (Go language).**

**Provide functions to calculate and judge the business days and holidays in each country.**

Currently, we only handle holidays for a few countries, but we plan to handle holidays for other countries in the future.  
Please Contribute!

## Install
```bash
$ go install github.com/yut-kt/goholiday
```

## Usage

- version >= v1.0.0
  - look at [example test file](https://github.com/yut-kt/goholiday/blob/master/goholiday_example_test.go) or [godoc example](https://godoc.org/github.com/yut-kt/goholiday)

- version = v0.x.x (deprecated)
  - look at [docs/v0.md](docs/v0.md)

## Supported ccTLD
```
support 2022, 2023
|- am (Armenia)
|- do (Dominican Republic)
|- gr (Greece)
|- jp (Japan)
|- sg (Singapore)
|- uk (United Kingdom)
    |- England
    |- NorthernIreland
    |- Scotland
    |- Wales
```

## Contribution
[CONTRIBUTING.md](docs/CONTRIBUTING.md)

## License
goholiday is released under the [MIT License](LICENSE).