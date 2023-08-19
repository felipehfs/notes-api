package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/felipehfs/gonotes/dtos"
	"github.com/felipehfs/gonotes/infra"
	"github.com/felipehfs/gonotes/repositories"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type UserController struct {
	Repository repositories.UserAbstractRepository
}

func NewUserController(repository repositories.UserAbstractRepository) UserController {
	return UserController{
		Repository: repository,
	}
}

func sendError(err []string, statusCode int, w http.ResponseWriter) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]any{
		"errors": err,
	})
}

func (u *UserController) Register(w http.ResponseWriter, r *http.Request) {
	var request dtos.CreateUser
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		sendError([]string{err.Error()}, http.StatusInternalServerError, w)
		return
	}

	if errors := infra.ValidateStruct(request); len(errors) > 0 {
		sendError(errors, http.StatusBadRequest, w)
		return
	}

	password, err := infra.HashPassword(request.Password)
	if err != nil {
		sendError([]string{err.Error()}, http.StatusInternalServerError, w)
		return
	}

	id, _ := uuid.NewRandom()

	err = u.Repository.Create(id.String(), request.Email, password)

	if err != nil {
		sendError([]string{err.Error()}, http.StatusBadRequest, w)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

func (u UserController) Login(w http.ResponseWriter, r *http.Request) {
	var request dtos.LoginUser
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		sendError([]string{err.Error()}, http.StatusBadRequest, w)
		return
	}

	if errors := infra.ValidateStruct(request); len(errors) > 0 {
		sendError(errors, http.StatusBadRequest, w)
		return
	}

	user, err := u.Repository.FindOne(request.Email)
	if err != nil {
		sendError([]string{err.Error()}, http.StatusBadRequest, w)
		return
	}

	if isValidPassword := infra.ComparePasswordHash(request.Password, user.Password); !isValidPassword {
		sendError([]string{infra.ErrorInvalidPassword.Error()}, http.StatusBadRequest, w)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"id":    user.Id,
		"email": user.Email,
		"token": infra.CreateToken(user.Id, user.Email, time.Minute*50),
	})
}

func (u *UserController) Routes() chi.Router {
	r := chi.NewRouter()
	r.Post("/register", u.Register)
	r.Post("/login", u.Login)
	return r
}
