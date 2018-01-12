package main

import "fmt"

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

func main(){
	var c stack
	fmt.Println(c)
	c.push(1,2,3,4,5,6,7,8,9,10)
	fmt.Println(c)
	n := c.pop()
	fmt.Println(n)
	fmt.Println(c)
	c.push(10,11)
	fmt.Println(c)

}