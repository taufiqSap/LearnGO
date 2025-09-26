package main

import "fmt"

func Array(){

/*	var buah [4]string
	buah[0] = "mangga"
	buah[1] = "apel"
	buah[2] = "jeruk"
	buah[3] = "semangka"

	fmt.Println(buah)

	buah2:= [4]string{"mangga", "apel", "jeruk", "semangka"}
	fmt.Println(buah2)

	*/
}

func Slice(){

	var buah = []string{"mangga", "apel", "jeruk", "semangka"}
	fmt.Println(buah[0:2]) // outpu mangga apel, karena index 2 tidak dihitung

	var aBuah = buah[0:3]
	var bBuah = buah[1:4]

	var Abuah = buah [1:2]
	var Bbuah = buah [0:1]

	fmt.Println(buah)
	fmt.Println(aBuah)
	fmt.Println(bBuah)
	fmt.Println(Abuah)
	fmt.Println(Bbuah)

	Bbuah[0] = "nanas"

	fmt.Println(buah)
	fmt.Println(aBuah)
	fmt.Println(bBuah)
	fmt.Println(Abuah)
	fmt.Println(Bbuah)
	fmt.Println(Bbuah)
	
}