package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main (){
	reader:=bufio.NewReader(os.Stdin)

	fmt.Println("enter any rating between 1 to 5")
	input,_:=reader.ReadString('\n')

	fmt.Println("thanks for your rating of ",input)


	numrating,err:=strconv.ParseFloat(strings.TrimSpace(input),64)
	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println("enhanced rating by 1 :",numrating+1)
	}
	

	
}