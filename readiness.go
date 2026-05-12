package main

import (
	"net/http"
)


func handlerReadiness(resw http.ResponseWriter, req *http.Request) {
		resw.Header().Add("Content-Type","text/plain; charset=utf-8")
		resw.WriteHeader(http.StatusOK)
		resw.Write([]byte("OK"))	
		// or resw.Write([]byte(http.StatusText(http.StatusOK)))
}

