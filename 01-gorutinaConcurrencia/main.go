package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Printf("Numero de CPU al inicio: %v\n", runtime.NumCPU())
	fmt.Printf("Numero de Gorutinas al inicio: %v\n", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("Hola desde la primera go routina.")
		wg.Done()
	}()
	go func() {
		fmt.Println("Hola desde la segunda go routina.")
		wg.Done()
	}()
	fmt.Printf("Numero de CPU en el medio: %v\n", runtime.NumCPU())
	fmt.Printf("Numero de Gorutinas en el medio: %v\n", runtime.NumGoroutine())

	wg.Wait()
	fmt.Println("A punto de finalizar main")

	fmt.Printf("Numero de CPU al final: %v\n", runtime.NumCPU())
	fmt.Printf("Numero de Gorutinas al final: %v\n", runtime.NumGoroutine())

}
