package main

import "fmt"

const englishGreetingsPrefix = "Hello, "
const spanishGreetingsPrefix = "¡Hola, "
const frenchGreetingsPrefix = "Bonjour, "
const koreanGreetingsPrefix = "안녕, "
const greetingsSuffix = "!"

// Hello is a public function
func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return getPrefix(language) + name + greetingsSuffix
}

// getPrefix is a private function
func getPrefix(language string) (prefix string) {
	// the named output variable does a couple of things
	// 1. defines a variable `prefix`, sets it equal to the "zero value" of the type
	// 2. returns it if there is a `return ` statement without anything specified
	// 3. will be displayed in the go docs

	switch language {
		case "Korean": 
			prefix = koreanGreetingsPrefix
		case "Spanish": 
			prefix = spanishGreetingsPrefix
		case "French": 
			prefix = frenchGreetingsPrefix
		default:
			prefix = englishGreetingsPrefix
	}	

	return
}

func main() {
	fmt.Println(Hello("world", "English"))
}