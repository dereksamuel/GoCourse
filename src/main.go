package main

import (
	"fmt"

	c "github.com/dereksamuel/gocourse/src/car"
	"github.com/google/uuid"
)

func main() {
	var otherCar = c.CarPublic{Label: "sadsadsdsd", Owner: c.Person{Name: "Jorge", Lastname: "Paul"}}

	fmt.Printf("Hello, world %s\n", "👁️")
	fmt.Println(uuid.New().String())
	fmt.Println(otherCar)
}
