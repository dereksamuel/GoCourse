package main

import "fmt"

func msgMaker(text string, channel chan string) {
	channel <- text
}

func main() {
	channel := make(chan string, 2)

	channel <- "Hola"
	// channel <- "Mundo"

	fmt.Println(len(channel), cap(channel)) // length and capacity

	// para cerrar el canal cuando no lo estes usando mas
	close(channel)

	// TODO: run this with an error:
	// channel <- "Mundo3"

	for message := range channel {
		fmt.Println(message)
	}

	// Select
	email1 := make(chan string)
	email2 := make(chan string)

	go msgMaker("Msg 1", email1)
	go msgMaker("Msg 2", email2)

	for i := 0; i < 2; i++ {
		select {
		case ml1 := <-email1:
			fmt.Println("Send email", ml1)
		case ml2 := <-email2:
			fmt.Println("Send email", ml2)
		}
	}
}
