package main 

import (
	"net/http"
	"encoding/json"
	"log"
)

func validateChirp( resw http.ResponseWriter, req *http.Request) {
	type responseVal struct {
		Valid bool `json:"valid"`
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
	respondWithJSON(resw, http.StatusOK, responseVal{Valid: true})

}

