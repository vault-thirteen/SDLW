package c

import (
	"fmt"
	"strings"
)

const ErrBadWord = "bad word: %s"

type Word string

func (w Word) Check() (err error) {
	if strings.Contains(string(w), " ") {
		return fmt.Errorf(ErrBadWord, string(w))
	}

	return nil
}

// StringWithoutWords removes specified words from a string.
func StringWithoutWords(sIn string, words []Word) (sOut string) {
	sOut = sIn
	for _, w := range words {
		sOut = strings.ReplaceAll(sOut, string(w), "")
	}
	return sOut
}
