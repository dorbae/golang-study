package main

import (
	"fmt"

	"github.com/dorbae/golang-study/dictionary/dictionary"
)

func main() {
	// #2.4 Ditionary part One
	// https://nomadcoders.co/go-for-beginners/lectures/1516
	dictionary := dictionary.Dictionary{"first": "First word"}
	dictionary["key1"] = "value1"
	fmt.Println(dictionary)

	definition, err := dictionary.Search("second")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

	// #2.5 Add method
	// https://nomadcoders.co/go-for-beginners/lectures/1517
	word := "key2"
	def := "value2"
	dictionary.Add(word, def)
	if err != nil {
		fmt.Println(err)
	}
	foundDefinition, err := dictionary.Search(word)
	fmt.Println("found", word, "definition:", foundDefinition)
	err2 := dictionary.Add(word, def)
	if err2 != nil {
		fmt.Println(err2)
	}

	// #2.6 Update Delete
	// https://nomadcoders.co/go-for-beginners/lectures/1518
	baseWord := "key3"
	dictionary.Add(baseWord, "value3")
	err3 := dictionary.Update(baseWord, "newValue3")
	if err3 != nil {
		fmt.Println(err)
	}
	definition3, _ := dictionary.Search(baseWord)
	fmt.Println(definition3)

	baseWord4 := "key4"
	dictionary.Add(baseWord4, "value4")
	dictionary.Delete(baseWord4)
	foundDefinition4, err4 := dictionary.Search(baseWord4)
	if err4 != nil {
		fmt.Println(err4)
	} else {
		fmt.Println(foundDefinition4)
	}
}
