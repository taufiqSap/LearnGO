package main

import (
	"fmt"
	"math/rand"
	"time"
)
/*func Point() {

	 rand.Seed(time.Now().UnixNano())

	 point := rand.Intn(100)+1

	 fmt.Println("your point score:", point)

	if point == 100{
		fmt.Println("you have a perfect score")
	} else if point  >= 91 && point <= 100{
		fmt.Println(" you have an A")
	} else if point >= 81 && point <=90{
		fmt.Println(" you have a B")
	} else if point >= 70&& point <=80{
		fmt.Println(" you have a C")
	} else{
		fmt.Println("you have failed")
	}
} */


/* func Usia(){

	var usia int = 0

	fmt.Print("Masukkan Usia Anda: ")
	fmt.Scan(&usia)
	
	switch{
	case usia >= 0 && usia <=5:
		fmt.Print("anda balita")
	case usia >5 && usia <= 15:
		fmt.Print("anda anak-anak")
	case usia >15 && usia <= 25:
		fmt.Print("anda remaja")
	case usia >25 && usia <= 45:
		fmt.Print("anda dewasa")
	case usia >45 && usia <= 65:
		fmt.Print("anda sudah tua")
	case usia >65:
		fmt.Print("anda lansia")
	default:
		fmt.Print("Anda bukan manusia")
	} 
} */

 func Usia() {

	rand.Seed(time.Now().UnixNano())
	usia := rand.Intn(100)+1

	fmt.Println("Usia Anda: ",usia)

	if usia >= 0 && usia <=5 {
		fmt.Print("anda balita")
	} else if usia >5 && usia <= 15 {
		fmt.Print("anda anak-anak")
	} else if usia >15 && usia <= 25 {
		fmt.Print("Anda remaja")
	} else if usia >25 && usia <= 45 {
		fmt.Print("anda dewasa")
	} else if usia >45 && usia <= 65 {
		fmt.Print("anda sudah tua")
	} else if usia >65 {
		fmt.Print("anda lansia")
	} else {
		fmt.Print("Anda bukan manusia")
	}
}




