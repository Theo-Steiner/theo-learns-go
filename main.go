package main

import (
	di "learngowithtests/di"
	"log"
	"net/http"
	"os"
)

func main() {
	di.Greet(os.Stdout, "Elodie")
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(di.MyGreeterHandler)))
}
