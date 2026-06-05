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
	platform string
	jwtSecret	string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	platform := os.Getenv("PLATFORM")
	if platform == "" {
		log.Fatal("PLATFORM must be set")
	}

	dbURL := os.Getenv("DB_URL")
	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error opening database: %s",err)
	}

	secret := os.Getenv("JWTSecret")
	if secret == "" {
		log.Fatal("cannot access JWTSecret")
	}

	dbQueries := database.New(dbConn)

	const filePath = "."
	const port = "8080"
	apiCfg := apiConfig{
		fileserverHits: atomic.Int32{},// i didnt did this 
		db: dbQueries,
		platform: platform,
		jwtSecret:	secret,
	}

	mux := http.NewServeMux()
	
	fsHandler := apiCfg.middlewareMetricsInc(http.FileServer(http.Dir(filePath)))
	mux.Handle("/app/", http.StripPrefix("/app",fsHandler))

	mux.HandleFunc("GET /api/healthz", handlerReadiness)
	mux.HandleFunc("GET /admin/metrics", apiCfg.handlerCount)
	mux.HandleFunc("POST /admin/reset", apiCfg.handlerReset)
	mux.HandleFunc("POST /api/chirps", apiCfg.handlerChirps)
	mux.HandleFunc("POST /api/users",apiCfg.handlerUser)
	mux.HandleFunc("GET /api/chirps", apiCfg.handlerGetChirps)
	mux.HandleFunc("GET /api/chirps/{chirpID}", apiCfg.handlerGetChirpByID)
	mux.HandleFunc("POST /api/login", apiCfg.handlerLogin)
	mux.HandleFunc("POST /api/refresh",apiCfg.handlerRefresh)
	mux.HandleFunc("POST /api/revoke", apiCfg.handlerRevoke)
	mux.HandleFunc("PUT /api/users", apiCfg.handlerUpdateUser)
	mux.HandleFunc("DELETE /api/chirps/{chirpID}", apiCfg.handlerDeleteChirp)

	s := &http.Server {
			Addr:	":" + port,
			Handler:	mux,
	}


	log.Printf("Serving files from %s on port: %s\n", filePath, port)
	log.Fatal(s.ListenAndServe())

}

