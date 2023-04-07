package main

import(
	"fmt"
)

func main(){
	var Int int
	Int = -12
	fmt.Println("number:",Int)

	var Uint uint
	Uint = 12
	fmt.Println("uint", Uint)

	var String string
	String = "My variable string"
	fmt.Println("string", String)

	var Boolean bool
	Boolean  = true
	fmt.Println("Boolean",Boolean)	

	fmt.Println("My variable address  is:", &Boolean)
	fmt.Println("My variables address is:", &Int)
	fmt.Println("My variables address is:", &String)
	fmt.Println("My variable address is:",&Uint)

	myInt := 12
	myString := "cristopher"
	fmt.Println("My variable is:",myInt, myString)

	const Const = "a12"
	fmt.Println("Constante",Const)

	const Integer int = 12
	fmt.Println("Constante Entero",Integer)

	fmt.Println()
	var my8bitsIntVar int8
	fmt.Printf("Int default value: %d\n", my8bitsIntVar)
}