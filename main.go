package main

import (
	"dictToAnki/anki"
	"dictToAnki/file_parse"
	"fmt"
)

func main() {
	words, err := file_parse.ParseWordsFromHtmlFile("./data/eudict2.html")
	if err != nil {
		panic(err)
	}

	client := anki.NewConnectClient("http://localhost:8765")
	decks, err := client.GetDecksNames()
	fmt.Println(decks, err)
	for _, word := range words {
		fmt.Println("create note: ", word.Word)
		id, err := client.AddNote("Default", word.Word, word.Meaning, []string{"test"})
		fmt.Println(id, err)
	}

}
