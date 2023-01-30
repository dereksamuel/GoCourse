package main

import (
	"fmt"
)

type Pc struct {
	Ram       int
	Disk_rang int
	Brand     string
}

// un ping nos va a retornar un mensaje a detalle
func (myPC Pc) PingMethod(anotherMsg string) {
	fmt.Println("-*-*-*-*-*-*-*-*-* ping -*-*-*-*-*-*-*-*-*")
	fmt.Println(anotherMsg)
	fmt.Println(myPC.Brand, "Pong")
	fmt.Println("-*-*-*-*-*-*-*-*-* ping end -*-*-*-*-*-*-*-*-*")
}

func (myPC *Pc) DuplicateRam() {
	myPC.Ram = myPC.Ram * 2
}

// *pc: acceder a values del valor en memoria de myPC

func (myPCArg Pc) String() string {
	return fmt.Sprintf("Ram: %d, DiskRange %d, Brand %s", myPCArg.Ram, myPCArg.Disk_rang, myPCArg.Brand) // va a crear un string
}

func main() {
	myPc := Pc{Ram: 15, Brand: "msi", Disk_rang: 100}

	fmt.Println(myPc.String())
	fmt.Println(myPc)
}
