package main 

import (
		"net/http"
		

)

func (cfg *apiConfig) handlerGetChirps(resw http.ResponseWriter, req *http.Request) {
		receivedChirps ,err := cfg.db.RetrieveChirps(req.Context())
		if err != nil {
			respondWithError(resw, http.StatusInternalServerError, "could request chirps", err)
			return 
		}
		
		chirps := []respondChirp{}
		
		for _, dbChirp := range receivedChirps {
				c := respondChirp{
					ID:		dbChirp.ID,
					CreatedAt:		dbChirp.CreatedAt,
					UpdatedAt: 		dbChirp.UpdatedAt,
					Body:					dbChirp.Body,
					UserID:				dbChirp.UserID,
			}
			chirps = append(chirps, c)
		}
		respondWithJSON(resw, http.StatusOK, chirps)

	
}
