package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	welcome:="namaskaram"
	fmt.Println(welcome)

	reader :=bufio.NewReader(os.Stdin)
    fmt.Println("please enter your name  ")

	//COMMMA OK || ERR OK
	input,_:=reader.ReadString('\n')
	fmt.Println("Thanks for coming  ",input )
	fmt.Printf("type of this input is %T",input)
}
