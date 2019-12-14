package main

import "fmt"

const hello = "Hello, "

// Hello ... utils func
func Hello(str string) string {
	return hello + str
}

func main() {
	fmt.Println(Hello("Chris"))
}
