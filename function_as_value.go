package main 

import "fmt"

func getGoodbye(name string) string{
	return "Good Bye " + name
}

func numer(a int) int{
	return a + 1
}

func main(){
	
	result := getGoodbye
	fmt.Println(result("bro"))

	result2 := numer
	fmt.Println(result2(1000))
}
