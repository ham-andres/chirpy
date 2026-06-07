package main 

import (
		"net/http"
		"github.com/google/uuid"
		

)

func (cfg *apiConfig) handlerGetChirps(resw http.ResponseWriter, req *http.Request) {
		chirps := []respondChirp{}

		author := req.URL.Query().Get("author_id")


		if author != "" {
				authorID, err := uuid.Parse(author)
				if err != nil {
						respondWithError(resw, http.StatusInternalServerError, "couldnt parse author id", err)
						return
				}
				chirpsByAuthor, err := cfg.db.RetrieveChirpsByAuthor(req.Context(), authorID)
				if err != nil {
						respondWithError(resw, http.StatusNotFound,"couldnt find chirp by author", err)
						return
				}
		
				for _, authorChirps := range chirpsByAuthor {
					c_author := respondChirp {
						ID:						authorChirps.ID,
						CreatedAt:		authorChirps.CreatedAt,
						UpdatedAt:		authorChirps.UpdatedAt,
						Body:					authorChirps.Body,
						UserID:				authorChirps.UserID,
					}
					chirps = append(chirps, c_author)
				}
			
				respondWithJSON(resw, http.StatusOK, chirps)

		} else {
			// without author returning all 
				receivedChirps ,err := cfg.db.RetrieveChirps(req.Context())
				if err != nil {
						respondWithError(resw, http.StatusInternalServerError, "could request chirps", err)
						return 
				}
		
				
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
		return

}
