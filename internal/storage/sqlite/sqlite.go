package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/satyamkanungo-dev/student-api/internal/config"
	"github.com/satyamkanungo-dev/student-api/internal/types"
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

func (s *SqliteStorage) GetStudentById(id int64) (types.Student, error) {
	stmt, err := s.DB.Prepare("SELECT * FROM students WHERE id = ? LIMIT 1")
	if err != nil {
		return types.Student{}, err
	}
	defer stmt.Close()

	var student types.Student

	// the order == database order
	if err = stmt.QueryRow(id).Scan(&student.ID, &student.Name, &student.Email, &student.Age); err != nil {
		// if  row is invalid or not there
		if err == sql.ErrNoRows {
			return types.Student{}, fmt.Errorf("no student found with id: %d", id)
		}

		// general error
		return types.Student{}, fmt.Errorf("query error: %w", err)
	}

	return student, nil
}
