package main
import "fmt"
func main(){
	n := []int{3,2,8,6,4,3}
	fmt.Println(Max(n))
}


func Max(a []int) (m int) {
	for _,v := range(a) {
		if v>m{
			m = v	
		}
	}
	return
}