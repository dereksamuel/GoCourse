package main

import (
	"fmt"

	"github.com/dereksamuel/gocourse/src/pc"
)

func main() {
	a := 50
	b := &a // va a ser la direccion en memoria donde se guarda a

	fmt.Println(b)
	fmt.Println(*b) // el valor al que le apunta la direccion en memoria

	*b = 780
	fmt.Println(a) // a nos va a dar 780 porque esta apuntando a la misma direccion en memoria de b

	myPc := pc.Pc{Ram: 16, Disk_rang: 82.5, Brand: "MSI"}
	myPc.PingMethod("Hello new msg")

	fmt.Println(myPc.Ram)
	myPc.DuplicateRam()
	fmt.Println(myPc.Ram)

	myPc.DuplicateRam()
	fmt.Println(myPc.Ram)
}
