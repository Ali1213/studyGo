package main

import (
	"fmt"
	"strconv"
)

type stack struct {
	i int
	data [10]int
}

func (a *stack) push(args ...int){
	for _, v := range(args){
		a.data[a.i] = v
		a.i++
	}
}

func (a *stack) pop()(c int){
	a.i--
	c = a.data[a.i]
	a.data[a.i] = 0
	// remove(a.data[a.i])
	return 
}

func (a *stack) String()(s string){
	for i:=0;i<a.i;i++{
		s+= "[" + strconv.Itoa(i) + ":" + strconv.Itoa(a.data[i])+"]"
	}
	return
}

func main(){
	var c stack
	c.push(1,2,3,4,5,6,7,8,9,10)
	fmt.Printf("My stack %v\n",&c)
}