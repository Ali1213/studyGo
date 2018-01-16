package main
import "fmt"
func main(){
	p:=plusTwo()
	fmt.Println(p(2))

}

func plusTwo()(func(int)int){
		f := func(a int)int{
			return a+2
		}
		return f
}