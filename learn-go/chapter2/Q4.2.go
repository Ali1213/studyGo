package main
import (
	"fmt"
	"unicode/utf8"
)
func main(){
	s:="asSASA ddd dsjkdsjs dk"
	runeLen := utf8.RuneCountInString(s)
	fmt.Println("runes =",runeLen)
	fmt.Println("bytes =",len(s))
}