package main

import (
	"fmt"

	"github.com/boutros/marc"
)

func createEmptyMARC(upc string) *marc.Record {
	m := marc.NewRecord()
	df := marc.NewDField("020")
	df = df.AddSubField("a", upc)
	m.AddDField(df)
	return m
}

func setFieldMARC(r *marc.Record, tag, sf, value string) {
	//upc := r.GetDFields("020")[0].SubFields[0].Value

	df := marc.NewDField(tag)
	df = df.AddSubField(sf, value)
	r.AddDField(df)
}

func aaaa() {
	r := createEmptyMARC("000000")
	setNameMARC(r, "hello world")
	fmt.Printf("%q", r)
}
