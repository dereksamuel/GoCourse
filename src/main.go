package main

import "fmt"

type figures2d interface { // para verificar las funciones
	getArea() float64 // solo funciona porque los structs tienen el method getArea
}

type quad struct {
	base float64
}

func (newQuad quad) getArea() float64 {
	return newQuad.base * newQuad.base
}

type rectangle struct {
	base   float64
	height float64
}

func (newRectangle rectangle) getArea() float64 {
	return newRectangle.base * newRectangle.height
}

func calculate(figures figures2d) {
	fmt.Println("Area:", figures.getArea())
}

func main() {
	mySquad := quad{base: 5}

	myRectangle := rectangle{base: 5, height: 2}

	calculate(mySquad)
	calculate(myRectangle)

	// lista de interfaces
	myInterface := []interface{}{"Hola", 1, 2, 3, 84, -5}
	fmt.Println(myInterface)
}
