package repositories

import (
	"database/sql"

	"github.com/felipehfs/gonotes/models"
)

type NoteRepository struct {
	Db *sql.DB
}

func NewNoteRepository(db *sql.DB) NoteRepository {
	return NoteRepository{
		Db: db,
	}
}

func (n *NoteRepository) Create(note models.Note) error {
	tx, err := n.Db.Begin()
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	_, err = tx.Exec("INSERT INTO notes (id, name, description, ownerId) VALUES (?, ?, ?, ?)",
		note.Id, note.Name, note.Description, note.OwnerId)

	if err != nil {
		return err
	}

	return tx.Commit()
}

func (n *NoteRepository) FindById(id string) (*models.Note, error) {
	var note models.Note
	err := n.Db.QueryRow("SELECT id, name, description, ownerId  FROM notes WHERE id=?", id).Scan(
		&note.Id,
		&note.Name,
		&note.Description,
		&note.OwnerId,
	)
	if err != nil {
		return nil, err
	}

	return &note, nil
}

func (n *NoteRepository) Read() (notes []models.Note, err error) {

	rows, err := n.Db.Query("SELECT id, name, description, created_at, ownerId FROM notes")
	if err != nil {
		return
	}

	var note models.Note
	for rows.Next() {
		rows.Scan(&note.Id, &note.Name, &note.Description, &note.CreatedAt, &note.OwnerId)
		notes = append(notes, note)
	}

	return
}
