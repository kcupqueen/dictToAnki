package main

import (
	"dictToAnki/anki"
	"dictToAnki/file_parse"
	log "github.com/sirupsen/logrus"
)

func main() {
	words, err := file_parse.ParseWordsFromHtmlFile("E:/Game/En101/eudict2023.html")
	if err != nil {
		panic(err)
	}

	client := anki.NewConnectClient("http://localhost:8765")
	decks, err := client.GetDecksNames()
	if err != nil {
		panic(err)
	}
	log.Infof("avaliable decks are: %v", decks)
	for _, word := range words {
		// log.Info("create note: ", word.Word)
		id, err := client.AddNote("Default", word.Word, word.Meaning, []string{"test"})
		log.Info("create note: ", word.Word, " id: ", id, " err: ", err)
	}

}
