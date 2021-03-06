package main

import "fmt"

const englishGreetingsPrefix = "Hello, "
const spanishGreetingsPrefix = "¡Hola, "
const frenchGreetingsPrefix = "Bonjour, "
const koreanGreetingsPrefix = "안녕, "
const greetingsSuffix = "!"

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	prefix := englishGreetingsPrefix

	switch language {
		case "Korean": 
			prefix = koreanGreetingsPrefix
		case "Spanish": 
			prefix = spanishGreetingsPrefix
		case "French": 
			prefix = frenchGreetingsPrefix
	}

	return prefix + name + greetingsSuffix
}

func main() {
	fmt.Println(Hello("world", "English"))
}