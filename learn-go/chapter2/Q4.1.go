package	 main	

import "fmt"

func main(){
	var s string
	extend := "A"
	for i:= 0; i < 100;i++{
		s+=extend
		fmt.Println(s);
	}
}