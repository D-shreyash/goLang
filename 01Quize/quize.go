package main

import "fmt"

func main(){

	var variable [10]int;
	
	for i:=0;i<10;i++{
		variable[i]=10;
	}
	fmt.Println(variable)
}