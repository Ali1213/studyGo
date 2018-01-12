package main

import "fmt"

func main(){

}

func fibner(n int)(int){
	totals := make([]int{1,1})

	if n < len(totals) {
		totals = append(totals,totals(len(totals)-1)+totals(len(totals)-2))
	}
	
	return totals(len(totals)-1)

}