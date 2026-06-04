package main 

import (
		"net/http"
		"log"
		"encoding/json"
		"time"

		"github.com/ham-andres/chirpy/internal/auth"
		"github.com/ham-andres/chirpy/internal/database"
)

func (cfg *apiConfig) handlerLogin(resw http.ResponseWriter, req *http.Request) {
		type parameters struct {
			Email string `json:"email"`
			Password string `json:"password"`
			ExpiresIn	int	`json:"expires_in_seconds"`
		}

		type responseVal struct {
			User
			Token string `json:"token"`
			RefreshToken	string	`json:"refresh_token"`
		}
		decoder := json.NewDecoder(req.Body)
		userParams := parameters{}

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
		if err != nil || !flag {
			log.Printf("error while verifying password %v", err)
			respondWithError(resw, http.StatusInternalServerError, "Incorrect email or password", err)
			return
		}

	// duration decision 
		expirationAccessToken := time.Hour
		expirationRefreshToken := time.Now().Add(60 * 24 * time.Hour)

		accessToken, err := auth.MakeJWT(logUser.ID, cfg.jwtSecret, expirationAccessToken) // duration and jwtsecret link them
		if err != nil {
			respondWithError(resw, http.StatusInternalServerError, "couldn't create token", err)
			return
		}
		
		refreshToken := auth.MakeRefreshToken()
		_, err = cfg.db.CreateToken(req.Context(), database.CreateTokenParams{
				Token:			refreshToken,
				UserID:			logUser.ID,
				ExpiresAt:	expirationRefreshToken,
		})
		if err != nil {
				respondWithError(resw, http.StatusInternalServerError, "couldn't store the refresh token", err)
				return
		}

		respondWithJSON(resw, http.StatusOK, responseVal{
			User:	User{
					ID:		logUser.ID,
					CreatedAt: 		logUser.CreatedAt,
					UpdatedAt:		logUser.UpdatedAt,
					Email:				logUser.Email,
			},
			Token:		accessToken,
			RefreshToken:	refreshToken,
			

		})	
}
