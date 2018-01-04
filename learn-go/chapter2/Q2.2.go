package main
import "fmt"

func main(){
	i:=0
	TAG:
	 fmt.Println(i)
	 i++
	 if i==10{
		 return
	 }
	 goto TAG
}