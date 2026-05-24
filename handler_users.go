package main 

import (
	"net/http"
	"encoding/json"
	"log"
	"time"

	"github.com/ham-andres/chirpy/internal/auth"
	"github.com/ham-andres/chirpy/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig)handlerUser(resw http.ResponseWriter, req *http.Request)  {
	type user struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}

	type responseVal struct {
		ID uuid.UUID `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Email string `json:"email"`
	}

	decoder := json.NewDecoder(req.Body)
	userParams := user{}

	err := decoder.Decode(&userParams)
	if err != nil {
		log.Printf("error while decoding %v", err)
		respondWithError(resw, http.StatusBadRequest, "couldn't decode user mail", err)

		return 
	}
	
	hashP, err := auth.HashPassword(userParams.Password)
	if err != nil {
		log.Printf("error while hashing :%v", err)
		respondWithError(resw, http.StatusBadRequest, "couldn't hash the password", err)
		return
	}

	params := database.CreateUserParams{
				Email:	userParams.Email,
				HashedPassword:	hashP,
	}

	createdUser, err := 	cfg.db.CreateUser(req.Context(), params)
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
