package testexp

import "fmt"

func MyFunc(input *string) {
	fmt.Println("INPUT", input)
	panic("panic panic")
	fmt.Println("Panic done")
}
