package main

import "fmt"

func Loop(){

for i :=1; i<=5; i++{
	for j:=1; j<=i; j++{
		fmt.Print("*")
	}
	fmt.Println()
}

}