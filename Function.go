package main

import (
	"fmt"
	//"strings"
	"math/rand"
	"time"
)

var randomizer = rand.New(rand.NewSource(time.Now().Unix()))

func String(){

	//var nama = []string{"Jhon","Doe"}
	//printMesagge("halo", nama)

	var randomValue  int

	randomValue = randomWithRange(1,100)
	fmt.Println("random Number ", randomValue)
	
	randomValue = randomWithRange(2,10)
	fmt.Println("random Number ", randomValue)

	randomValue = randomWithRange(2, 10)
    fmt.Println("random number:", randomValue)
}

func randomWithRange(min, max int) int {
	var value = randomizer.Int()%(max-min+1) + min
	return value
}
/* func printMesagge(message string, arr[]string){
	var namaString = strings.Join(arr," ")
	fmt.Println(message, namaString)
} */