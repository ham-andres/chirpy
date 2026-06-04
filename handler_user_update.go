package main 

import (
		"net/http"
		"encoding/json"

		
		"github.com/ham-andres/chirpy/internal/auth"
		"github.com/ham-andres/chirpy/internal/database"
)

func (cfg *apiConfig) handlerUpdateUser (resw http.ResponseWriter,  req *http.Request) {
	type parameters struct {
			Email 			string	`json:"email"`
			Password		string	`json:"password"`
	}

	type responseVal struct {
			User
	}

	accessToken, err := auth.GetBearerToken(req.Header)
	if err != nil {
			respondWithError(resw, http.StatusUnauthorized, "couldn't access the Access Token", err)
			return
}
	
	decoder := json.NewDecoder(req.Body)
	params := parameters{}

	err = decoder.Decode(&params)
	if err != nil {
			respondWithError(resw, http.StatusBadRequest,"Couldn't decode json", err)
			return
	}
	
	userId, err := auth.ValidateJWT(accessToken, cfg.jwtSecret)
	if err != nil {
			respondWithError(resw, http.StatusUnauthorized,"Incorrect token or tokenSecret", err)
			return
	}

	// hash the password before sending to database
	hashedPassword, err := auth.HashPassword(params.Password)
	if err != nil {
			respondWithError(resw, http.StatusBadRequest, "couldn't hash the password", err)
			return
	}
	
	user, err := cfg.db.UpdateUserEmailAndPassw(req.Context(), database.UpdateUserEmailAndPasswParams{
				Email:			params.Email,
				HashedPassword:		hashedPassword,
				ID:					userId,
	})
	if err != nil {
			respondWithError(resw, http.StatusUnauthorized,"couldn't update email and password", err)
			return
	}
	respondWithJSON(resw, http.StatusOK, User{
		ID:					user.ID,
		CreatedAt:	user.CreatedAt,
		UpdatedAt:	user.UpdatedAt,
		Email:			user.Email,
	})

}
