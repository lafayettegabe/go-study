package main

import "fmt"

const englishHelloPrefix = "Hello, "

var HelloPrefix = map[string]string{
	"en": "Hello, ",
	"es": "Hola, ",
	"fr": "Bonjour, ",
	"pt": "Olá, ",
	"jp": "こんにちは, ", // Konnichiwa
	"kr": "안녕하세요, ", // Annyeonghaseyo
	"cn": "你好, ",    // Nǐ hǎo
}

func HelloWorld() string {
	return "Hello World!"
}

func HelloYou(name string) string {
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("%s%s!", englishHelloPrefix, name)
}

func HelloYouLang(name, lang string) string {
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("%s%s!", HelloPrefix[lang], name)
}

func main() {
	fmt.Println(HelloWorld())
}
