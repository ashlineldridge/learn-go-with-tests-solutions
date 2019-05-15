package main

import "fmt"

const (
	defaultLang = "English"
	defaultName = "World"
)

var greetings = map[string]string{
	"English": "Hello, ",
	"Spanish": "Hola, ",
	"French": "Bonjour, ",
}


func Hello(name string, lang string) string {
	if name == "" {
		name = defaultName
	}

	if lang == "" {
		lang = defaultLang
	}

	prefix := greetings[lang]

	if prefix == "" {
		prefix = greetings[defaultLang]
	}

    return prefix + name
}

func main() {
    fmt.Println(Hello("", ""))
}
