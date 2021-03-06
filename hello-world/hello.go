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

	if language == "Korean" {
		return koreanGreetingsPrefix + name + greetingsSuffix
	}

	if language == "Spanish" {
		return spanishGreetingsPrefix + name + greetingsSuffix
	}

	if language == "French" {
		return frenchGreetingsPrefix + name + greetingsSuffix
	}

	return englishGreetingsPrefix + name + greetingsSuffix
}

func main() {
	fmt.Println(Hello("world", "English"))
}