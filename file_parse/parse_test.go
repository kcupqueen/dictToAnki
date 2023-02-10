package file_parse

import (
	"fmt"
	"testing"
)

func TestParseWordsFromHtmlFile(t *testing.T) {
	words, err := ParseWordsFromHtmlFile("./test.html")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(len(words))
}
