package file_parse

import (
	"encoding/csv"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

// ParseWordsFromCsvFile Cambridge Dictionary CHINESE-ENGLISH CSV
func ParseWordsFromCsvFile(path string) ([]Word, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = file.Close()
	}()

	var words []Word
	reader := csv.NewReader(file)
	var index = 0
	for {
		index++
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		if index == 1 {
			continue
		}
		if hasEmptyInStrings(line[:3]) {

			log.Warnf("line has empty: %v", line[1])
			continue
		}

		words = append(words, Word{Word: line[1], Meaning: line[3]})

	}
	return words, nil
}

func hasEmptyInStrings(ss []string) bool {
	for _, s := range ss {
		if s == "" {
			return true
		}
	}
	return false
}
