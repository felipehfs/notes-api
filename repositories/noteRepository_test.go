package repositories_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/felipehfs/gonotes/models"
	"github.com/felipehfs/gonotes/repositories"
	"github.com/google/uuid"
)

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("an error %v was not expected", err)
	}

	repository := repositories.NoteRepository{
		Db: db,
	}

	id, err := uuid.NewRandom()
	if err != nil {
		t.Errorf("an error occurred generating id: %v", err)
	}

	onwerId, err := uuid.NewRandom()
	if err != nil {
		t.Errorf("an error occurred generating onwer id: %v", err)
	}

	note := models.Note{
		Id:          id.String(),
		Name:        "Compras a prazo",
		Description: "lorem ipsum",
		OwnerId:     onwerId.String(),
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO notes").WithArgs(note.Id, note.Name, note.Description, note.OwnerId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = repository.Create(note)
	if err != nil {
		t.Error(err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were expectations not fulfullied %v", err)
	}
}

func TestReadNotes(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("an error %v was not expected", err)
	}

	repository := repositories.NewNoteRepository(db)

	rows := sqlmock.NewRows([]string{"id", "name", "description", "created_at", "ownerId"}).
		AddRow("120201", "teste", "teste", "2020-12-28 10:00:00", "120102012")

	mock.ExpectQuery("SELECT id, name, description, created_at, ownerId FROM notes").WillReturnRows(rows)

	_, err = repository.Read()
	if err != nil {
		t.Errorf("an unexpected error occurred: %v", err)
	}

}
