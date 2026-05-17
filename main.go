package main	

import (
	"net/http"
	"log"
	"sync/atomic"
	"os"
	"database/sql"

	"github.com/ham-andres/chirpy/internal/database"
  _ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

type apiConfig struct {
	fileserverHits atomic.Int32
	db *database.Queries
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)

	const filePath = "."
	const port = "8080"
	apiCfg := apiConfig{
		fileserverHits: atomic.Int32{},// i didnt did this 
		db: dbQueries,
	}

	mux := http.NewServeMux()
	
	fsHandler := apiCfg.middlewareMetricsInc(http.FileServer(http.Dir(filePath)))
	mux.Handle("/app/", http.StripPrefix("/app",fsHandler))

	mux.HandleFunc("GET /api/healthz", handlerReadiness)
	mux.HandleFunc("GET /admin/metrics", apiCfg.handlerCount)
	mux.HandleFunc("POST /admin/reset", apiCfg.handlerReset)
	mux.HandleFunc("POST /api/validate_chirp", validateChirp)	

	s := &http.Server {
			Addr:	":" + port,
			Handler:	mux,
	}


	log.Printf("Serving files from %s on port: %s\n", filePath, port)
	log.Fatal(s.ListenAndServe())

}

