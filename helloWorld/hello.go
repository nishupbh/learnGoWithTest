package main

import "fmt"

/*
1.	Write a test
2.	Make the compiler pass
3.	Run the test, see that it fails and check the error message is meaningful
4.	Write enough code to make the test pass
5. Refactor*/

//Spanish for Constant name
const Spanish = "spanish"

//French for Constant name
const French = "french"

// EnglishHelloPrefix for english greeting prefix
const EnglishHelloPrefix = "Hello, "

// SpanishHelloPrefix for english greeting prefix
const SpanishHelloPrefix = "Hola, "

// FrenchHelloPrefix for english greeting prefix
const FrenchHelloPrefix = "Bonjour, "

func hello(name string, language string) string {

	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case Spanish:
		prefix = SpanishHelloPrefix
	case French:
		prefix = FrenchHelloPrefix
	default:
		prefix = EnglishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(hello("Nishant", "french"))
}
