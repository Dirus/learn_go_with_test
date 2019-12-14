package main

import "fmt"

// Hello ... utils func
func Hello(str string) string {
	return "Hello " + str
}

func main() {
	fmt.Println(Hello("Chris"))
}
