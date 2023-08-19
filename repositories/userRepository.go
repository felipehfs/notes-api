package repositories

import (
	"database/sql"

	"github.com/felipehfs/gonotes/models"
)

type UserRepository struct {
	Db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return UserRepository{
		Db: db,
	}
}

func (u UserRepository) FindOne(email string) (*models.User, error) {
	var user models.User
	err := u.Db.QueryRow(`
		SELECT id, email, password FROM users WHERE email=?
	`, email).Scan(&user.Id, &user.Email, &user.Password)

	if err != nil {
		return nil, err
	}

	return &user, err
}

func (u UserRepository) Create(id string, email string, password string) (err error) {
	tx, err := u.Db.Begin()
	if err != nil {
		return
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	if _, err = tx.Exec("INSERT INTO users(id, email, password, active) VALUES (?, ?, ?, ?)", id, email, password, true); err != nil {
		return
	}

	return
}
