package main

import	(
		"net/http"
	  "github.com/google/uuid"
		"github.com/ham-andres/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerDeleteChirp(resw http.ResponseWriter, req *http.Request) {
	accessToken, err := auth.GetBearerToken(req.Header)
	if err != nil {
			respondWithError(resw, http.StatusUnauthorized, "Couldn't get token", err)
			return
	}

	verifiedUserID, err := auth.ValidateJWT(accessToken, cfg.jwtSecret)
	if err != nil {
			respondWithError(resw, http.StatusForbidden, "couldn't validate token", err)
			return
	}

	chirpId, err := uuid.Parse(req.PathValue("chirpID"))
	if err != nil {
			respondWithError(resw, http.StatusNotFound, "No chirp belong to id", err)
			return
	}

// wrong  delete chirp only by id not user id too, use the user id to verify the chirp is his or not
	chirp, err := cfg.db.GetChirpByID(req.Context(), chirpId)
	if err != nil {
			respondWithError(resw, http.StatusNotFound,"no chirp by this chirpID", err)
			return
	}
	if chirp.UserID != verifiedUserID {
			respondWithError(resw, http.StatusForbidden,"Not authorized to delete chirp", err)
			return
	}

	err = cfg.db.DeleteChirp(req.Context(), chirp.ID)
	if err != nil {
			respondWithError(resw, http.StatusNotFound,"not authorized to delete", err)
			return 
	}

	resw.WriteHeader(http.StatusNoContent)
	return
}
