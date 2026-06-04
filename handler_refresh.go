package main 

import (
		"net/http"
		"time"

		"github.com/ham-andres/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerRefresh(resw http.ResponseWriter, req *http.Request) {

	type response struct {
			Token string `json:"token"`
	}

	token, err := auth.GetBearerToken(req.Header)
	if err != nil {
			respondWithError(resw, http.StatusUnauthorized, "couldn't get token", err)
			return 
	}
	user, err := cfg.db.GetUserFromRefreshToken(req.Context(), token)
	if err != nil {
			respondWithError(resw, http.StatusUnauthorized, "invalid token", err)
			return 
	}
	
	refreshToken, err := auth.MakeJWT(user.UserID, cfg.jwtSecret, time.Hour)
	if err != nil {
			respondWithError(resw, http.StatusUnauthorized, "couldn't refresh token", err)
			return
	}

	respondWithJSON(resw, http.StatusOK,response{
		Token:		refreshToken,
	})
}
