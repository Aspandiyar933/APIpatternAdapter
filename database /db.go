package database 

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

type MySQLStorage struct {
	db *sql.DB
}

func NewMySQLStorage(cfg mysql.Config) *MySQLStorage {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MySQL!")
	return &MySQLStorage{
		db: db,
	}
}

func (s *MySQLStorage) Init() (*sql.DB, error) {
	if err := s.createTaskTable(); err != nil {
		return nil, err
	}
}

func (s *MySQLStorage) createTaskTable() error {
	_, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS Todo (
			id INT UNSIGNED NOT NULL AUTO_INCREMENT,
			description TEXT NOT NULL
		)
	`)
	return err 
}