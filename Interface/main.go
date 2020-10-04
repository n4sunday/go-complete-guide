package main

import (
	"fmt"
	"net/http"
	"os"
)

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	fmt.Println(resp)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	return "Hi There!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}
