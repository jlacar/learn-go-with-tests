package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const hebrew = "Hebrew"

const englishPrefix = "Hello, "
const frenchPrefix = "Bonjour, "
const hebrewPrefix = "Shalom, "
const spanishPrefix = "Hola, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return prefix(language) + name + "!"
}

func prefix(language string) (prefix string) {
	switch (language) {
	case french:
		prefix = frenchPrefix
	case hebrew:
		prefix = hebrewPrefix
	case spanish:
		prefix = spanishPrefix
	default:
		prefix = englishPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
