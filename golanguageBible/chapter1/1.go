package main
import (
	"strings"
	"os"
	"fmt"
)
func main(){
	var s string
	// for i:=0;i<len(os.Args);i++{
	// 	s += sep + os.Args[i]
	// 	sep = " "
	// }
	s = strings.Join(os.Args[1:]," ")
	fmt.Println(s);
}