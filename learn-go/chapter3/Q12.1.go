package main
import "fmt"

func main(){
	a := []int{1,2,3,4,5,6}

	f:= Map(d,a)
	fmt.Println(f)
}

func d(c int) int{
	fmt.Println(c)
	return c
}

func Map(f func(int)int,a []int)([]int){
	var n = make([]int,len(a))
	for i,v := range(a){
		n[i] = f(v)
	}
	return n
}