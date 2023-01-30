package main

import (
	"fmt"
	"sync"
)

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done() // esperar que sea la ultima que se ejecute

	fmt.Println(text)
}

func main() {
	var wg sync.WaitGroup // es mas complejo mantenerlo

	fmt.Println("Hello")
	wg.Add(1)              // anadimos una goroutine
	go say("Hello 2", &wg) // ejecutar con concurrencia y no queda en mismo hilo de ejecucion de main que acaba en siguiente linea
	// time.Sleep(time.Second * 1) // no es recomendado, pero le da tiempo necesario a la goroutine
	fmt.Println("Hello 3")
	wg.Wait() // esperar hasta que wg.Done se ejecute

	// funcion anonima
	go func(msg string) {
		fmt.Println(msg)
	}("Hello")
}
