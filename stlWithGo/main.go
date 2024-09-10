package main

import (
	"fmt"
	"stl/internal"
)

func main() {
	fmt.Println("main package initialized")
	//stack initialization
	stack := internal.NewStack(5)
	queue := internal.NewQueue()
	fmt.Println("Stack  initialized", stack)
	fmt.Println("Queue  initialized", queue)

}
