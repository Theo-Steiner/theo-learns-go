package main

import (
	di "learngowithtests/di"
	"os"
)

func main() {
	di.Greet(os.Stdout, "Elodie")
}
