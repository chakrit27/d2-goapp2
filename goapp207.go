package main

import "fmt"

func main() {
	var a float64 = 90.0 //Static Type
	a = 50.5
	var b string = "Welcome";

	y := 14 //Dynamic Type

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(y)
	fmt.Printf("A is of Type %T\n",a)
	fmt.Printf("A is of Type %T",y)
}