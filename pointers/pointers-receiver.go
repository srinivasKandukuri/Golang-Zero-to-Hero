package main

import "fmt"

type Point struct {
	X, Y int
}

// function that takes a pointer
func Scale(p *Point, factor int) {
	p.X = p.X * factor
	p.Y = p.Y * factor
}

func Print(p *Point) {
	fmt.Println("X: ", p.X, ", Y: ", p.Y)
}

// method with a receiver pointer
func (p *Point) Scale(factor int) {
	p.X = p.X * factor
	p.Y = p.Y * factor
}

func (p *Point) Print() {
	fmt.Println("X: ", p.X, ", Y: ", p.Y)
}

func (p Point) ScaleVal(factor int) {
	p.X = p.X * factor
	p.Y = p.Y * factor
}

func (p Point) PrintP() {
	fmt.Println("X: ", p.X, ", Y: ", p.Y)
}

func main() {
	//Use case 1: both uses same pointer type No implicit conversion
	p := Point{X: 2, Y: 3}

	p_ptr := &p
	Scale(p_ptr, 2)
	Print(p_ptr)

	fmt.Println("Pointer receiver case 1")
	p_ptr.Scale(2)
	p_ptr.Print()

	// receiver arg is of type Point, compiler implicitly get &p to match type with receiver param
	fmt.Println("Pointer receiver case 2")
	p.Scale(2) // compiler implicitly get &p
	p.Print()

	// receiver arg is of type Point, compiler implicitly get &p to match type with receiver param
	// case 3: only receiver argument is pointer type
	fmt.Println("Pointer receiver case 3")
	p_ptr.ScaleVal(2)
	p_ptr.PrintP()
}
