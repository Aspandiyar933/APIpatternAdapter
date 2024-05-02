package store

import (
	"database/sql"

	"github.com/Aspandiyar933/APIpatternAdapter/pattern"
)

type Store interface {
	CreateTask(t *pattern.Todo) (*pattern.Todo)
	GetTask(id int64) (*pattern.Todo) 
}

type Storage struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) CreateTask(t *pattern.Todo) (*pattern.Todo, error) {
	rows, err := s.db.Exec("INSERT INTO Todo (description, do) VALUES(?, ?)",
		t.Description, t.Do,
	)
	if err != nil {
		return nil, err
	}
	id, err := rows.LastInsertId()
	if err != nil {
		return nil, err
	}
	t.ID = id
	return t, nil
}

func (s *Storage) GetTask(id int64) (*pattern.Todo, error) {
	var t pattern.Todo
	err := s.db.QueryRow("SELECT id, description, do FROM Todo WHERE id = ?", id).Scan(
		&t.ID, &t.Description, &t.Do)
	return &t, err
}