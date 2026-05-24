package main 

import (
		"net/http"
		"log"
		"encoding/json"
		"time"

		"github.com/ham-andres/chirpy/internal/auth"
		"github.com/google/uuid"
)

func (cfg *apiConfig) handlerLogin(resw http.ResponseWriter, req *http.Request) {
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
			respondWithError(resw, http.StatusBadRequest, "couldn't decode json", err)
			return
		}

		logUser, err := cfg.db.GetUserByEmail(req.Context(), userParams.Email)
		if err != nil {
			log.Printf("error while loggin with email %v", err)
			respondWithError(resw, http.StatusBadRequest, "couldn't verify with email", err)
			return
		}

		flag, err := auth.CheckPasswordHash(userParams.Password, logUser.HashedPassword)	
		if err != nil {
			log.Printf("error while verifying password %v", err)
			respondWithError(resw, http.StatusInternalServerError, "couldn't verify password", err)
			return
		}

		if !flag {
			respondWithError(resw, http.StatusUnauthorized, "Incorrect email or password", err)
			return
		} else {

		}
		respondWithJSON(resw, http.StatusOK, responseVal{
			ID:		logUser.ID,
			CreatedAt: 		logUser.CreatedAt,
			UpdatedAt:		logUser.UpdatedAt,
			Email:				logUser.Email,
		})	
}
