package main
import "fmt"

type e interface{}


func main(){
	a := []e{1,2,3,4,5,6}
	b := []e{"a","b","c","d"}
	f:= Map(d,a)
	e:= Map(d,b)
	fmt.Println(f)
	fmt.Println(e)
}

func d(f e) e{
	switch f.(type){
	case int:
		return f.(int) * 2
	case string:
		return f.(string) + f.(string) + f.(string) + f.(string)
	}

	return f
}

func Map(f func(e)e,n []e)([]e){
	var c = make([]e,len(n))
	for i,v := range(n){
		c[i] = f(v)
	}
	return c
}