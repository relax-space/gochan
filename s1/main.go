package main

import "fmt"

//
func main() {
	//read()
	// deadlock()
	// deadlock2()
	unlock()
}

func read() {
	ch := make(chan int)
	go func() {
		fmt.Println(1)
		v := <-ch
		fmt.Println(v)
	}()
	ch <- 2
	fmt.Println("3")
}

func deadlock() {
	ch := make(chan int)
	ch <- 1
	fmt.Println("1")
}

func deadlock2() {
	ch := make(chan int)
	a := <-ch
	fmt.Println(a)
}

func unlock() {
	ch := make(chan int, 1)
	ch <- 1
	a := <-ch
	fmt.Println(a)
}
