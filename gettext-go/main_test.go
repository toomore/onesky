package main

import (
	"os"
	"testing"
)

func TestReadCSV(*testing.T) {
	readCSV("onesky.csv")
}
func TestCreatePO(*testing.T) {
	csvdata, _ := readCSV("onesky.csv")
	createPO("onesky.po", csvdata, 0, "test_temp")
	os.RemoveAll("./test_temp")
}

func TestCSVtopo(*testing.T) {
	csvtopo("onesky.csv", "./test_temp")
}

func BenchmarkReadCSV(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		readCSV("onesky.csv")
	}
}
