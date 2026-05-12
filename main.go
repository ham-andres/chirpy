package main	

import (
	"net/http"
	"log"
)

func main() {
	const filePath = "."
	const port = "8080"

	mux := http.NewServeMux()
	s := &http.Server {
		Addr:	":" + port,
		Handler:	mux,
	}

	mux.Handle("/app/", http.StripPrefix("/app/", http.FileServer(http.Dir(filePath))))
	mux.HandleFunc("/healthz", handlerReadiness)

	log.Printf("Serving files from %s on port: %s\n", filePath, port)
	log.Fatal(s.ListenAndServe())

}

func handlerReadiness(resw http.ResponseWriter, req *http.Request) {
		resw.Header().Set("Content-Type","text/plain; charset=utf-8")
		resw.WriteHeader(http.StatusOK)
		resw.Write([]byte("OK"))	
}
