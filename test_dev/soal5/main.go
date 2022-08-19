package main

import "fmt"

type Kendaraan struct {
	Suara string
}

func (k Kendaraan) Akselerasi() {
	k.Suara = "Swoosh"
	fmt.Println(k.Suara)
}

type Sepeda struct {
	Kendaraan
	rantai string
}




func main() {
	var s Sepeda
	s.Kendaraan.Akselerasi()
}

