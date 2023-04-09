package main

import "fmt"

const helloEnglishPrefix = "Hello, "
const helloSpanishPrefix = "Hola, "
const helloFrenchPrefix = "Bonjour, "
const helloItalianPrefix = "Ciao, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return generatePrefix(language) + name
}

func generatePrefix(language string) (prefix string) {
	switch language {
	case "Spanish":
		prefix = helloSpanishPrefix
	case "French":
		prefix = helloFrenchPrefix
	case "Italian":
		prefix = helloItalianPrefix
	default:
		prefix = helloEnglishPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("Martin", "Spanish"))
}
