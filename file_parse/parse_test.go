package file_parse

import (
	"fmt"
	"testing"
)

func TestParseWordsFromHtmlFile(t *testing.T) {
	words, err := ParseWordsFromHtmlFile("./eudict_cambridge.html")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(len(words))
}

func TestParseWordsFromCsvFile(t *testing.T) {
	words, err := ParseWordsFromCsvFile("./test_csv.csv")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(len(words))
}
