package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/boutros/marc"
)

func createEmptyMARC(upc string) *marc.Record {
	m := marc.NewRecord()
	setCFieldMARC(m, "005", time.Now().Format("20060102150405.0"))
	setDFieldMARC(m, "020", "a", upc)
	return m
}

func setCFieldMARC(r *marc.Record, tag, value string) {
	c := marc.CField{
		Tag:   tag,
		Value: value,
	}
	r.SetCField(c)
}

func setDFieldMARC(r *marc.Record, tag, sf, value string) {
	df := marc.NewDField(tag)
	df = df.AddSubField(sf, value)
	r.AddDField(df)
}

func setTypeMARC(r *marc.Record, tstr string) {
	setCFieldMARC(r, "006", tstr)
}

func setNameMARC(r *marc.Record, name string) {
	setDFieldMARC(r, "222", "a", name)
}

func saveMARC(r *marc.Record) error {
	upc := r.GetDFields("020")[0].SubField("a")
	f, err := os.OpenFile(fmt.Sprintf("records/%s.marc", upc), os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	e := marc.NewEncoder(f, marc.MARC)
	er := e.Encode(r)
	if er != nil {
		return er
	}

	if err := f.Close(); err != nil {
		return err
	}
	return nil
}

func getMARCRecord(upc string) (*marc.Record, error) {
	f, err := os.OpenFile(fmt.Sprintf("records/%s.marc", upc), os.O_RDONLY, 0755)
	if err != nil {
		return &marc.Record{}, err
	}
	d := marc.NewDecoder(f, marc.MARC)
	return d.Decode()
}

func getMARCRecords() (string, error) {
	f, err := os.Open("records")
	if err != nil {
		return "", err
	}
	s, err := f.Readdirnames(0)
	out := []string{}
	for _, item := range s {
		if item != ".gitignore" {
			out = append(out, item)
		}
	}
	var b []byte
	b, err = json.Marshal(out)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", b), nil
}

func aaaa() {
	r := createEmptyMARC("000000")
	setNameMARC(r, "hello world")
	fmt.Printf("%q", r)
}
