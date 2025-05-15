package main

import "fmt"

func main() {

	//type declaration of a variable 
	var myname string="Jaipreeth"
	fmt.Println(myname)
	fmt.Printf("type of the variable :%T\n",myname)
	
	//boolean variable 

	var ismale bool =true
	fmt.Println(ismale)
	fmt.Printf("type of the variable :%T\n",ismale)

	//int 
	var age int =21
	fmt.Println(age)
	fmt.Printf("type of the variable :%T\n",age)


	//small float 

	var height float32=5.111133344455
	fmt.Println(height)
	fmt.Printf("type of the variable :%T\n",height)

	//float 
	var heightt float64=5.111133344455
	fmt.Println(heightt)
	fmt.Printf("type of the variable :%T\n",heightt)

	//default values 

	var madness int 
	fmt.Println(madness)

	var respect string 
	fmt.Println(respect)

	//implicit type

	var drama =0
	var mom ="kusuma kumari"
	fmt.Println(drama)
	fmt.Printf("type of the variable :%T\n",drama)
	fmt.Println(mom)
	fmt.Printf("type of the variable :%T\n",mom)


	//no var style

	noofachievements:=0
	fmt.Println(noofachievements)
	fmt.Printf("type of the variable :%T\n",noofachievements)

	//note this style is allowed only inside the method 

}