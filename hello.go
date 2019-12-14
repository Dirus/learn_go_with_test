package main

import "fmt"

const hello = "Hello, "

// Hello ... utils func
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return hello + name
}

func main() {
	fmt.Println(Hello("Chris"))
}
