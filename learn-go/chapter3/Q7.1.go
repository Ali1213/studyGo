package main

import "fmt"


func main(){
	a,b := 2,7
	c,d := rightPos(a,b)
	fmt.Println(c,d)
}

func rightPos(a,b  int) (c,d int){
	if a>b {
		c,d = b,a
	}else{
		c,d = a,b
	}
	return
}