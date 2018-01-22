package main 

import (
	"fmt"
	"os/exec"
)

func main(){
	cmd := exec.Command("/bin/ls","-l")
	buf,err := cmd.Output()
	fmt.Println(buf)
	fmt.Println(err)
}