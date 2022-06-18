package hello

import (
	"fmt"
	"os"
	"time"

	"github.com/lucianboboc/learn-go-with-tests/maths"
)

const english = "English"
const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func main() {
	fmt.Println(Hello("Lucian", "English"))
	maths.SVGWriter(os.Stdout, time.Now())
}

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	if language == "" {
		language = english
	}

	return getPrefix(language) + name
}

func getPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
