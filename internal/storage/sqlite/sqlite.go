package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/satyamkanungo-dev/student-api/internal/config"
)

type SqliteStorage struct {
	DB *sql.DB
}

// create a table
func New(cfg *config.Config) (*SqliteStorage, error) {
	db, err := sql.Open("sqlite3", cfg.Storage_path)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		email TEXT,
		age INTEGER
		)`)

	if err != nil {
		return nil, err
	}

	return &SqliteStorage{
		DB: db,
	}, nil
}

func (s *SqliteStorage) CreateStudent(name, email string, age int) (int64, error) {
	stmt, err := s.DB.Prepare("INSERT INTO students (name, email, age) VALUES (?,?,?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(name, email, age)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
