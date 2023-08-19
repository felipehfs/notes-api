package dtos

type CreateNote struct {
	Name        string `json:"name" validate:"required,min=4"`
	Description string `json:"description" validate:"required"`
	OwnerID     string `json:"ownerId" validate:"required"`
}
