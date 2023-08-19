package main

import (
	"net/http"

	"github.com/felipehfs/gonotes/controllers"
	"github.com/felipehfs/gonotes/infra"
	"github.com/felipehfs/gonotes/repositories"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	db, err := infra.ConnectToDatabase()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userRepository := repositories.NewUserRepository(db)
	noteRepository := repositories.NewNoteRepository(db)

	userHandler := controllers.NewUserController(userRepository)
	notesHandler := controllers.NewNoteController(noteRepository)

	router.Mount("/users", userHandler.Routes())
	router.Mount("/notes", notesHandler.Routes())

	http.ListenAndServe(":3080", router)
}
