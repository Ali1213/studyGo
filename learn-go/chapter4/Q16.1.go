package stack

import (
	"fmt"
	"strconv"
)

type stack struct {
	i int
	data [10]int
}

func (a *stack) Push(args ...int){
	for _, v := range(args){
		a.data[a.i] = v
		a.i++
	}
}

func (a *stack) Pop()(c int){
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
