package main

import "fmt"
import (
	hello "learngowithtests/hello"
	integers "learngowithtests/integers"
)

func main() {
	fmt.Println(hello.Hello("world", ""))
	fmt.Printf("Adding 3 & 5 is: %d, yay! \n", integers.Add(3, 5))
}
