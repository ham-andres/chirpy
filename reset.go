package main

import (
	"net/http"
	"log"
)

func (cfg *apiConfig) handlerReset(resw http.ResponseWriter, req *http.Request) {
	if cfg.platform != "dev" {
		resw.WriteHeader(http.StatusForbidden)
		resw.Write([]byte("Reset is only allowed in dev environment"))
		return 
	}
	
	cfg.fileserverHits.Store(0)
	err := cfg.db.DeleteUsers(req.Context())
	if err != nil {
		log.Printf("error while resetting, deleting users: %v", err)
		respondWithError(resw,http.StatusInternalServerError, "couldn't delete users", err)
		return 
	}
	resw.WriteHeader(http.StatusOK)
	resw.Write([]byte("Hits reset to 0 and database reset to initial state"))
}
