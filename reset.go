package main

import (
	"net/http"
)

func (cfg *apiConfig) handlerReset(resw http.ResponseWriter, req *http.Request) {
	cfg.fileserverHits.Store(0)
	resw.WriteHeader(http.StatusOK)
	resw.Write([]byte("Hits reset to 0"))
}
