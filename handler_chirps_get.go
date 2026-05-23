package main 

import (
		"net/http"
		"github.com/google/uuid"
		

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

func (cfg *apiConfig) handlerGetChirpByID( resw http.ResponseWriter, req *http.Request) {
		chirpId, err := uuid.Parse(req.PathValue("chirpID"))
		if err != nil {
			respondWithError(resw, http.StatusBadRequest, "Invalid ID", err)
			return
		}

		receivedChirp, err := cfg.db.GetChirpByID(req.Context(), chirpId)
		if err != nil {
			respondWithError(resw, http.StatusNotFound," Couldn't found chirp of given ID", err)
			return 
		}

		chirp := respondChirp{
			ID:		receivedChirp.ID,
			CreatedAt: 		receivedChirp.CreatedAt,
			UpdatedAt:		receivedChirp.UpdatedAt,
			Body:					receivedChirp.Body,
			UserID:				receivedChirp.UserID,
		}

		respondWithJSON(resw, http.StatusOK, chirp)

}
