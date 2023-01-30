package pc

import "fmt"

type Pc struct {
	Ram       int
	Disk_rang float32
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
