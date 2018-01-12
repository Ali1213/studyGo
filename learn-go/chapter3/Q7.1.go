package main

import "fmt"


func main(){
	a,b := 2,7
	c,d := rightPos(a,b)
	fmt.Println(c,d)
}

func rightPos(a,b  int) (int, int){
	if a>b {
		return b,a
	}
	return a,b
}