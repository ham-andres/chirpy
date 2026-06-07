package main 

import (
		"net/http"
		"errors"
		"database/sql"
		"encoding/json"
		"github.com/google/uuid"	
		"github.com/ham-andres/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerWebhook( resw http.ResponseWriter, req *http.Request) {
		type parameters struct {
				Event		string		`json:"event"`
				Data		struct {
						UserID		uuid.UUID		`json:"user_id"`
					}		`json:"data"`
		}
		
		apiKey, err := auth.GetAPIKey(req.Header)
		if err != nil {
				respondWithError(resw, http.StatusUnauthorized,"Couldn't get the api key", err)
				return
		}
		if apiKey != cfg.polkaKey {
				resw.WriteHeader(http.StatusUnauthorized)
				return
		}

		decoder := json.NewDecoder(req.Body)
		webhookParam := parameters{}

		err = decoder.Decode(&webhookParam)
		if err != nil {
				respondWithError(resw, http.StatusBadRequest, "couldn't decode the request json", err)
				return
		}

		if webhookParam.Event != "user.upgraded" {
			resw.WriteHeader(http.StatusNoContent)	
			return 
		}

		_, err = cfg.db.Upgrade(req.Context(),webhookParam.Data.UserID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows){
					respondWithError(resw, http.StatusNotFound,"couldn't find user", err)
					return
			}
				respondWithError(resw, http.StatusInternalServerError,"couldn't upgrade user", err)
				return 
		}
			
		resw.WriteHeader(http.StatusNoContent)
		return
}
