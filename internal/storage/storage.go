package storage

import "github.com/satyamkanungo-dev/student-api/internal/types"

type IStorage interface {
	CreateStudent(name, email string, age int) (int64, error)
	GetStudentById(id int64) (types.Student, error)
	GetStudents() ([]types.Student, error)
}
