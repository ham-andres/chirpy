package main 

import (
	"net/http"
	"encoding/json"
	"log"
	"strings"
)

func validateChirp( resw http.ResponseWriter, req *http.Request) {
	type responseVal struct {
		CleanedBody string `json:"cleaned_body"`
	}

	type parameters struct {
		Body string `json:"body"`
	}

	decoder := json.NewDecoder(req.Body)
	params := parameters {}
	err := decoder.Decode(&params)
	if err != nil {
		log.Println("error while decoding")
		respondWithError(resw, http.StatusInternalServerError, "couldn't decode parameters", err)
		return
	}
	const maxChirpLength = 140
	if len(params.Body) > maxChirpLength {
		respondWithError(resw, http.StatusBadRequest, "Chirp is too long", nil)
		return 
	}
	badWords := map[string]struct{}{
		"kerfuffle": {},
		"sharbert": {},
		"fornax": {},
	}
	cleanedBody := cleaner(params.Body, badWords)
	respondWithJSON(resw, http.StatusOK, responseVal{CleanedBody: cleanedBody})

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
