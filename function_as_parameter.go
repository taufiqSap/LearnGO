package main 

import (
	"fmt"
)

func helloWithFilter(name string, filter func(string) string){
	filteredName :=filter(name)
	fmt.Println("woy", filteredName)
}

func spamFilter(name string) string{
	if name == "goblok"{
		return "****"
	} else {
		return name
	}
}

func main(){

	helloWithFilter("mas", spamFilter)

	filter := spamFilter 
	helloWithFilter("goblok",filter)
}