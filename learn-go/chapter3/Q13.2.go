package main
import "fmt"
func main(){
	n := []int{3,2,8,6,4,3}
	fmt.Println(Min(n))
}


func Min(a []int) (m int) {
	if len(a) > 0 {
		m = a[0]
	}
	for _,v := range(a) {
		
		if v<m{
			m = v	
		}
	}
	return
}