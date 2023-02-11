package file_parse

import (
	"github.com/anaskhan96/soup"
	"golang.org/x/net/html"
	"os"
	"strings"
)

type Word struct {
	Word    string `json:"word"`
	Meaning string `json:"meaning"`
}

func ParseWordsFromHtmlFile(htmlFilePath string) ([]Word, error) {
	b, err := os.ReadFile(htmlFilePath)
	if err != nil {
		panic(err)
	}
	words := make([]Word, 0)
	doc := soup.HTMLParse(string(b))
	// println(doc.HTML())
	items := doc.FindAll("tr")
	for _, item := range items {
		tds := item.FindAll("td")
		var word, meaning string
		for i, td := range tds {
			if i == 1 {
				word = td.Text()
			}
			if i == 4 {
				div := td.Find("div")
				aTags := div.FindAll("a")
				for _, a := range aTags {
					for k, v := range a.Attrs() {

						if k == "href" {
							if strings.HasPrefix(v, "sound:") {
								a.Pointer.Attr = []html.Attribute{}
							}
						}
						imgs := a.FindAll("img")
						for _, img := range imgs {
							for k, v := range img.Attrs() {
								if k == "src" {
									if strings.HasPrefix(v, "file:") {
										//fmt.Println("delete img: ", v)
										img.Pointer.Attr = []html.Attribute{}
									}
								}
							}
						}
					}

				}
				meaning = div.HTML()
			}

		}
		if word != "" && meaning != "" {
			words = append(words, Word{
				Word:    word,
				Meaning: meaning,
			})

		}

	}
	return words, nil
}
