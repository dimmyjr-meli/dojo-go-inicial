package main

import "testing"
import "fmt"

func TestHello(t *testing.T) {
	result := Hello( "Jarvis")

	fmt.Println(result)
	if result != "Hello Jarvis" {
		t.Errorf("Deu ruim!!")
	}

}
