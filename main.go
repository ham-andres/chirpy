package main	

import (
	"net/http"
	"log"
	"sync/atomic"
)

type apiConfig struct {
	fileserverHits atomic.Int32
}

func main() {
	const filePath = "."
	const port = "8080"
	apiCfg := apiConfig{
		fileserverHits: atomic.Int32{}, // i didnt did this 
	}

	mux := http.NewServeMux()
	
	fsHandler := apiCfg.middlewareMetricsInc(http.FileServer(http.Dir(filePath)))
	mux.Handle("/app/", http.StripPrefix("/app",fsHandler))

	mux.HandleFunc("GET /api/healthz", handlerReadiness)
	mux.HandleFunc("GET /admin/metrics", apiCfg.handlerCount)
	mux.HandleFunc("POST /admin/reset", apiCfg.handlerReset)
	
	s := &http.Server {
			Addr:	":" + port,
			Handler:	mux,
	}


	log.Printf("Serving files from %s on port: %s\n", filePath, port)
	log.Fatal(s.ListenAndServe())

}

