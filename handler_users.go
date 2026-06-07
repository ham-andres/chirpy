package main 
import (
	"net/http"
	"encoding/json"
	"time"

	"github.com/ham-andres/chirpy/internal/auth"
	"github.com/ham-andres/chirpy/internal/database"
	"github.com/google/uuid"
)
type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	IsChirpyRed	bool		`json:"is_chirpy_red"`
}

func (cfg *apiConfig)handlerUser(resw http.ResponseWriter, req *http.Request)  {
	type parameters struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}

	type responseVal struct {
		User
	}

	decoder := json.NewDecoder(req.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {

		respondWithError(resw, http.StatusBadRequest, "couldn't decode user mail", err)

		return 
	}
	
	hashedPassword, err := auth.HashPassword(params.Password)
	if err != nil {
		respondWithError(resw, http.StatusInternalServerError, "Couldn't hash password", err)
		return 
	}

	createdUser, err := 	cfg.db.CreateUser(req.Context(), database.CreateUserParams{
		Email: 					params.Email,
		HashedPassword:	hashedPassword,
	})
	if err != nil {
		respondWithError(resw, http.StatusInternalServerError, "couldn't create user", err)
		return
	}
	
	respondWithJSON(resw, http.StatusCreated, responseVal{
			User: User{
					ID:						createdUser.ID,
					CreatedAt:		createdUser.CreatedAt,
					UpdatedAt:		createdUser.UpdatedAt,
					Email:				createdUser.Email,
					IsChirpyRed:	createdUser.IsChirpyRed,

			},
	})
	return

}
