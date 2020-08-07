package main

import (
	"fmt"

	"github.com/agatory/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	word := "hello"
	definition := "Greeting"
	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}
	hello, _ := dictionary.Search("hello")
	fmt.Println("found", word, "definition:", hello)
	err2 := dictionary.Add("word", "gree")
	if err2 != nil {
		fmt.Println(err2)
	}
	words, _ := dictionary.Search("word")
	fmt.Println(words)
}
