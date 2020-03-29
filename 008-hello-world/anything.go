package main

import "fmt"

func main() {

	fmt.Println("hello everyone")

	foo()

	fmt.Println("something else")

	for i := 0; i < 200; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	bar()

}
func foo() {
	fmt.Println("I'm in foo")
}

func bar() {
	fmt.Println("The end")
}
