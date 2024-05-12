package main

import "fmt"

func main () {
	fmt.Println(Hello("world", ""))
}

const(
  spanish = "Spanish"
  japanese = "Japanese"

  englishHelloPrefix = "Hello, "
  spanishHelloPrefix = "Hola, "
  japaneseHelloPrefix = "こんにちは"
)

func Hello (name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case japanese:
		prefix = japaneseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
