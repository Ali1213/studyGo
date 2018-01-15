package main
import "fmt"
func main(){
	p:=plusTwo(2)
	fmt.Println(p(2))

}

func plusTwo(n int)(func(int)int){
		f := func(a int)int{
			return a+n
		}
		return f
}