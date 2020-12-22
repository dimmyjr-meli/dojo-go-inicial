package main

import "fmt"

func main() {
	hi := Hello("Jarvis")
	fmt.Print(hi)
}


func Hello(s string) string {
	return "Hello " + s
}
