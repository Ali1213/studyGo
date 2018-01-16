package main

import "fmt"

func main(){
	fmt.Println(fibonacci(2))
	fmt.Println(fibonacci(5))
}

func fibonacci(n int)(int){
	totals := make([]int,2)
	totals[0] = 1
	totals[1] = 1

	for n > len(totals) {
		totals = append(totals,totals[len(totals)-1]+totals[len(totals)-2])
	}
	return totals[len(totals)-1]
}