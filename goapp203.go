package main

import "fmt"

var (
contry = "Thailand" //Global variable
province = "trat"
poscode = 23000
)

func main() {
	y:= true //Local variable
	fmt.Println(y)

	fmt.Println(contry)
	fmt.Println(province)
	fmt.Println(poscode)
}