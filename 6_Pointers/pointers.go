package main

import "fmt"

func main() {

	// declaring a pointer
	var p *int
	x := 10

	// Initialize the Pointer with the Address of x
	p = &x

	fmt.Println("value of x :", x)
	fmt.Println("Address of x :", &x)

	fmt.Println("value of p :", p)
	fmt.Println("Address of p :", *p)

	//Modify the Value at the Address the Pointer is Pointing To
	*p = 20

	fmt.Println("value of x :", x)
	fmt.Println("value of p :", *p)
	fmt.Println("address of  p :", p)
	fmt.Println("address of  x :", &x)
}
