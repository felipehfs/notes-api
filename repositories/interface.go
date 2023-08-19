package repositories

import "github.com/felipehfs/gonotes/models"

type UserAbstractRepository interface {
	FindOne(email string) (*models.User, error)
	Create(id string, email string, password string) error
}
