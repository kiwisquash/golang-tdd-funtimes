package main

import "fmt"

const englishGreetingsPrefix = "Hello, "
const greetingsSuffix = "!"

func Hello(name string) string {
	if name == "" {
		name = "world"
	}
	return englishGreetingsPrefix + name + greetingsSuffix
}

func main() {
	fmt.Println(Hello("world"))
}