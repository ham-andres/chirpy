package main 

import (
	"net/http"
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
)

func (cfg *apiConfig)handlerUser(resw http.ResponseWriter, req *http.Request)  {
	type user struct {
		Email string `json:"email"`
	}

	type responseVal struct {
		ID uuid.UUID `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Email string `json:"email"`
	}

	decoder := json.NewDecoder(req.Body)
	userMail := user{}

	err := decoder.Decode(&userMail)
	if err != nil {
		log.Printf("error while decoding %v", err)
		respondWithError(resw, http.StatusBadRequest, "couldn't decode user mail", err)

		return 
	}

	email := userMail.Email
	createdUser, err := 	cfg.db.CreateUser(req.Context(), email)
	if err != nil {
		log.Printf("error while creating user %v",err)
		respondWithError(resw, http.StatusInternalServerError, "couldn't create user", err)
		return
	}
	
	respondWithJSON(resw, http.StatusCreated, responseVal{
		ID:	createdUser.ID,
		CreatedAt:	createdUser.CreatedAt,
		UpdatedAt:	createdUser.UpdatedAt,
		Email:	createdUser.Email,
	})

}
