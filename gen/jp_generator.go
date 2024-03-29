package main

import (
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"go/format"
	"io"
	"log"
	"os"
	"text/template"
	"time"
)

type Holiday struct {
	Date string
	Name string
}

var filename = flag.String("output", "nholidays/jp/schedule.go", "output file name")

//go:generate go run . -output nholidays/jp/schedule.go
func main() {
	f, err := os.Open("nholidays/jp/national_holidays.csv")
	if err != nil {
		panic(err)
	}
	r := csv.NewReader(f)

	var (
		h      []*Holiday
		isSkip bool
	)

	for i := 0; ; i++ {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		if i == 0 {
			continue
		}

		var holiday Holiday
		for i, v := range record {
			switch i {
			case 0:
				t, err := time.Parse("2006/1/2", v)
				if err != nil {
					panic(err)
				}
				tmp := t.Format("2006-01-02")
				if tmp >= "2000-01-01" {
					holiday.Date = t.Format("2006-01-02")
					isSkip = false
				} else {
					isSkip = true
				}
			case 1:
				if !isSkip {
					holiday.Name = v
					h = append(h, &holiday)
				}
			}
		}
	}

	var v = struct {
		URL     string
		Holiday []*Holiday
	}{
		"https://www8.cao.go.jp/chosei/shukujitsu/syukujitsu.csv",
		h,
	}

	var buf bytes.Buffer
	err = template.Must(template.New("prog").Parse(prog)).Execute(&buf, v)
	if err != nil {
		log.Fatal(err)
	}
	data, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*filename)

	err = os.WriteFile(*filename, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

const prog = `
// Code generated by jp_generator.go; DO NOT EDIT.
// Based on holiday from {{.URL}}.

package jp

import (
	"time"

	"github.com/yut-kt/goholiday/nholidays"
)

func New() *nholidays.ScheduleImpl  {
	return nholidays.New(
		map[time.Weekday]struct{}{
			time.Saturday: {},
			time.Sunday:   {},
		},
		map[string]string{
			{{range .Holiday}}	"{{.Date}}": "{{.Name}}",
			{{end}}
		},
	)
}

`
