package main

import (
	"fmt"
	gochan "test-gochan"

	"github.com/chneau/limiter"
)

func main() {
	//forString()
	forPoint()
}

func forString() {
	limit := limiter.New(10)
	chars := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	for _, char := range chars {
		char := char //In order to ensure the correct data, this code can not be ignored
		limit.Execute(func() {
			fmt.Println(char)
		})
	}
	limit.Wait()
	fmt.Println("done")
}

func forPoint() {
	limit := limiter.New(10)
	fruits := []*gochan.Fruit{
		&gochan.Fruit{Name: "apple1"},
		&gochan.Fruit{Name: "apple2"},
		&gochan.Fruit{Name: "apple3"},
		&gochan.Fruit{Name: "apple4"},
		&gochan.Fruit{Name: "apple5"},
		&gochan.Fruit{Name: "apple6"},
		&gochan.Fruit{Name: "apple7"},
		&gochan.Fruit{Name: "apple8"},
		&gochan.Fruit{Name: "apple9"},
		&gochan.Fruit{Name: "apple10"},
		&gochan.Fruit{Name: "apple11"},
		&gochan.Fruit{Name: "apple12"},
		&gochan.Fruit{Name: "apple13"},
		&gochan.Fruit{Name: "apple14"},
		&gochan.Fruit{Name: "apple15"},
		&gochan.Fruit{Name: "apple16"},
		&gochan.Fruit{Name: "apple17"},
		&gochan.Fruit{Name: "apple18"},
		&gochan.Fruit{Name: "apple19"},
		&gochan.Fruit{Name: "apple20"},
		&gochan.Fruit{Name: "apple21"},
		&gochan.Fruit{Name: "apple22"},
		&gochan.Fruit{Name: "apple23"},
		&gochan.Fruit{Name: "apple24"},
		&gochan.Fruit{Name: "apple25"},
		&gochan.Fruit{Name: "apple26"},
		&gochan.Fruit{Name: "apple27"},
		&gochan.Fruit{Name: "apple28"},
		&gochan.Fruit{Name: "apple29"},
		&gochan.Fruit{Name: "apple30"},
	}
	for _, fruit := range fruits {
		fruit := fruit
		limit.Execute(func() {
			fmt.Println(fruit.Name)
		})
	}
	limit.Wait()
	fmt.Println("done")
}
