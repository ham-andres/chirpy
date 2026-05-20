package main 

import (
	"errors"
	"strings"
)

func validateChirps(body string) (string, error) {
	
	const maxChirpLength = 140
	if len(body) > maxChirpLength {
		return "", errors.New("chirp too long maxChirpLength = 140")
	}
	badWords := map[string]struct{}{
		"kerfuffle": {},
		"sharbert": {},
		"fornax": {},
	}
	cleanedBody := cleaner(body, badWords)
	return cleanedBody, nil

}


func cleaner (body string, badWords map[string]struct{}) string {
	words := strings.Split(body, " ")	
	for i, word := range words {
		loweredWord := strings.ToLower(word)
		if _, ok := badWords[loweredWord]; ok {
			words[i] = "****"
		}
	}
	cleaned := strings.Join(words, " ")
	return cleaned
}
