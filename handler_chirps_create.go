package main 

import (
	"net/http"
	"encoding/json"
	"log"
	"time"

  "github.com/google/uuid"
	"github.com/ham-andres/chirpy/internal/database"
	"github.com/ham-andres/chirpy/internal/auth"

)

	type respondChirp struct {
		ID        uuid.UUID		`json:"id"`
		CreatedAt time.Time		`json:"created_at"`
		UpdatedAt time.Time		`json:"updated_at"`
		Body      string			`json:"body"`
		UserID    uuid.UUID		`json:"user_id"`
	}



func (cfg *apiConfig) handlerChirps(resw http.ResponseWriter, req *http.Request) {
	type bodyFields struct {
		Body string `json:"body"` // removed userID after creating it from MakeJWT
	}
	
	token, err := auth.GetBearerToken(req.Header)
	if err != nil {
		respondWithError(resw, http.StatusUnauthorized, "Couldn't extract token", err)
		return
	}

  verifiedUserID, err := auth.ValidateJWT(token, cfg.jwtSecret)
	if err != nil {
		respondWithError(resw, http.StatusUnauthorized, "couldn't validate token", err)
		return 
	}


	// change is we validate before decoding this time, therefore above code of GetBearerToken
	decoder := json.NewDecoder(req.Body)
	bodyField := bodyFields{}
	err = decoder.Decode(&bodyField)
	if err != nil {
		log.Println("error while decoding body fields")
		respondWithError(resw, http.StatusBadRequest, "couldn't decode body field", err)
		return
	}
	cleanedBody, err := validateChirps(bodyField.Body)
	if err != nil {
		respondWithError(resw, http.StatusBadRequest, "couldn't validate chirp", err)
		return
	}

	

	chirpCreated, err := cfg.db.CreateChirp(req.Context(), database.CreateChirpParams{
		Body: cleanedBody,
		UserID: verifiedUserID,
	})

	
	if err != nil {
		log.Println("error while creating chirp in database")
		respondWithError(resw, http.StatusInternalServerError, "couldn't create chirp in database", err)
		return 
	}

	respondWithJSON(resw, http.StatusCreated, respondChirp{
		ID:						chirpCreated.ID,
		CreatedAt:		chirpCreated.CreatedAt,
		UpdatedAt:		chirpCreated.UpdatedAt,
		Body:					chirpCreated.Body,
		UserID:				chirpCreated.UserID,
	})
}


