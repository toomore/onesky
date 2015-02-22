package main

import "testing"

func TestReadCSV(*testing.T) {
	readCSV("onesky.csv")
}

func BenchmarkReadCSV(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		readCSV("onesky.csv")
	}
}
