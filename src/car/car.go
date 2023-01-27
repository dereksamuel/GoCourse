package car

import "fmt"

type Person struct {
	Name     string
	Lastname string
}

type CarPublic struct {
	Label string
	Owner Person
	year  int16
}

func init() {
	fmt.Printf("\nMy new car code is done\n")
}
