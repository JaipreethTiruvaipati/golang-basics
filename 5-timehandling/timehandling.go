package main

import (
	"fmt"
	"time"
)
func main(){
	fmt.Println("welcome to time study of golang" )


	presenttime:=time.Now()
	fmt.Println(presenttime)

	fmt.Println(presenttime.Format("01-02-2006 15:04:05 Monday"))

	createddate:=time.Date(2004,time.March,17,8,15,0,0,time.UTC)
	fmt.Println(createddate)
	fmt.Println(createddate.Format("01-02-2006 15:04:05 Monday"))
}