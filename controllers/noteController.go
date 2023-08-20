package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/felipehfs/gonotes/dtos"
	"github.com/felipehfs/gonotes/infra"
	"github.com/felipehfs/gonotes/models"
	"github.com/felipehfs/gonotes/repositories"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type NoteController struct {
	Repository repositories.NoteRepository
}

func NewNoteController(repository repositories.NoteRepository) NoteController {
	return NoteController{
		Repository: repository,
	}
}

func (n NoteController) Delete(w http.ResponseWriter, r *http.Request) {
	idParams := chi.URLParam(r, "id")
	if err := n.Repository.Delete(idParams); err != nil {
		sendError([]string{err.Error()}, http.StatusBadRequest, w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (n NoteController) Read(w http.ResponseWriter, r *http.Request) {
	notes, err := n.Repository.Read()
	if err != nil {
		sendError([]string{err.Error()}, http.StatusBadRequest, w)
		return
	}

	json.NewEncoder(w).Encode(notes)
}

func (n NoteController) Create(w http.ResponseWriter, r *http.Request) {
	var body dtos.CreateNote
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sendError([]string{err.Error()}, http.StatusBadRequest, w)
		return
	}

	if errors := infra.ValidateStruct(body); len(errors) > 0 {
		sendError(errors, http.StatusBadRequest, w)
		return
	}

	id, _ := uuid.NewRandom()

	err := n.Repository.Create(models.Note{
		Id:          id.String(),
		Name:        body.Name,
		Description: body.Description,
		OwnerId:     body.OwnerID,
	})

	if err != nil {
		sendError([]string{err.Error()}, http.StatusBadRequest, w)
		return
	}

	note, err := n.Repository.FindById(id.String())

	if err != nil {
		sendError([]string{err.Error()}, http.StatusBadRequest, w)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(note)
}

func (n NoteController) Routes() chi.Router {
	r := chi.NewRouter()

	r.Use(infra.JwtAuthenticate)
	r.Post("/", n.Create)
	r.Get("/", n.Read)
	r.Delete("/{id}", n.Delete)
	return r
}
