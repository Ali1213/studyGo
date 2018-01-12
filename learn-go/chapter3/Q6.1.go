package main

import "fmt"


func main(){
	s:=[]float64{1,2,3,4}
	a:= average(s)
	fmt.Println(a)
}

func average(a []float64) (r float64){
	var total float64

	if len(a) == 0 {
		r = 0
		return
	}

	for _, num := range(a){
		total += num
	}
	r = total / float64(len(a))
	return
}