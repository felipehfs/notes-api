package repositories_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/felipehfs/gonotes/repositories"
	"github.com/google/uuid"
)

func TestFindOneUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error %v was not expected", err)
	}

	email := "felipe@gmail.com"

	mock.ExpectQuery("^SELECT id, email, password FROM users WHERE email=?").WithArgs(email).
		WillReturnRows(sqlmock.NewRows([]string{`id`, `email`, `password`}).AddRow(1, email, "1234"))

	repository := repositories.UserRepository{
		Db: db,
	}

	_, err = repository.FindOne(email)
	if err != nil {
		t.Errorf("An error ocurred: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}

func TestRegisterUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error %v was not expected", err)
	}

	email := "felipe@gmail.com"
	password := "123456"

	repository := repositories.UserRepository{
		Db: db,
	}

	id, err := uuid.NewRandom()
	if err != nil {
		t.Errorf("an error occurred generating id: %v", err)
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO users").WithArgs(id, email, password, true).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	if err := repository.Create(id.String(), email, password); err != nil {
		t.Errorf("There were unfulfilled expectations: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}
