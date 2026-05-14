package main

import (
	"net/http"
	"log"
	"encoding/json"
)

func respondWithError(resw http.ResponseWriter, code int, msg string, err error) {
	if err != nil {
		log.Println(err)
	}
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}

	type errorResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(resw, code, errorResponse{
		Error: msg,
	})
}

func respondWithJSON(resw http.ResponseWriter, code int, payload interface{}) {
	resw.Header().Set("Content-Type", "application/json")
	response, err := json.Marshal(payload) 
	if err != nil {
		log.Printf("failed json marshalling: %s",err)
		resw.WriteHeader(500)
		return
	}

	resw.WriteHeader(code)
	resw.Write(response)
}
