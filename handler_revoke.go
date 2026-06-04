package main 

import (
		"net/http"
		"database/sql"
		"time"
		"github.com/ham-andres/chirpy/internal/database"
		"github.com/ham-andres/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerRevoke(resw http.ResponseWriter, req *http.Request) {
	refreshToken, err := auth.GetBearerToken(req.Header)
	if err != nil {
		respondWithError(resw, http.StatusBadRequest, "couln't get the token", err)
		return 
	}
	err = cfg.db.RevokeToken(req.Context(), database.RevokeTokenParams{
			UpdatedAt:		time.Now(),
			RevokedAt:		sql.NullTime{Time: time.Now(), Valid: true},
			Token:				refreshToken,

	})	 
	if err != nil {
		respondWithError(resw, http.StatusInternalServerError, "couldn't revoke the token", err)
		return
	}
	resw.WriteHeader(http.StatusNoContent)

}
