package storage

type IStorage interface {
	CreateStudent(name, email string, age int) (int64, error)
}
