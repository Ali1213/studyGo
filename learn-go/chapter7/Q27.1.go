package main

import "fmt"

var c chan int

func main(){
	c:= make(chan int)
	quit := make(chan bool)
	go shower(c,quit)
	for i:=0;i<10;i++ {
		c <- i
	}
}

func shower(c chan int, quit chan bool) {
	for{
		select{
		case j := <-c:
			fmt.Println(j)
		case <-quit:
			break
		}
	}
}