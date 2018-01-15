package main

import "fmt"

func main(){
	fmt.Println(fibner(10))
}

func fibner(n int)(int){
	totals := make([]int,2)
	totals[0] = 1
	totals[1] = 1

	for n > len(totals) {
		totals = append(totals,totals[len(totals)-1]+totals[len(totals)-2])
	}
	
	return totals[len(totals)-1]

}