package main
import "fmt"

func main(){
	a:= []int{2,4,5,8,1,7,6,9}
	fmt.Println(BubbleSort(a))
}

func BubbleSort(a []int) ([]int){

	for i:=0;i<len(a)-1;i++{
		for j:=i+1;j<len(a);j++{
			if a[i]>a[j] {
				a[i],a[j] = a[j],a[i]
			}
		}
	}
	return a
}