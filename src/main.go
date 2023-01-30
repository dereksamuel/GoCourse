package main

import "fmt"

func say(text string, c chan<- string) { // solo de entrada
	c <- text // se le manda el text atributte COM UN PIPELINE
}

func main() {
	channel := make(chan string, 1) // tipo del dato del canal, y buena practica es definirle cantidad limite

	go say("Bye", channel)
	fmt.Println("Hello")
	fmt.Println(<-channel)
	<-channel // ESTO BLOQUEA LA FUNC main
}
