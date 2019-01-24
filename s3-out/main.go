package main

import (
	"fmt"

	"github.com/chneau/limiter"
)

func main() {
	forString()
}

func forString() {
	limit := limiter.New(10)
	chars := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	outChan := make(chan string, len(chars))
	for _, char := range chars {
		char := char
		limit.Execute(func() {
			if char < "l" {
				outChan <- ""
				return
			}
			outChan <- char
		})
	}
	limit.Wait()
	close(outChan)
	for char := range outChan {
		if len(char) != 0 {
			fmt.Println(char)
		}
	}
	fmt.Println("done")
}
