package main

import "fmt"

func main(){
	fmt.Println("let us learn pointers")

	var  ptr *int
	fmt.Println("value of pointer is ",ptr)

	number:=17
	var myptr *int=&number
	fmt.Println("pointer address of number ",myptr)
	fmt.Println("value at  address of number ",*myptr)

	*myptr =*myptr+2

	fmt.Println("the updated value of number ",*myptr)

}