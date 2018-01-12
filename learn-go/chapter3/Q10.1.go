package main

import "fmt"


func main(){
	printArgs(2,3,4,5)
}

func printArgs(args ...int){
	for _,v := range(args){
		fmt.Println(v)
	}
}