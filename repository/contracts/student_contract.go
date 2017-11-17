package contract

import (
	"ajar-go/models"
	"database/sql"
)

type StudentRepository interface {
	Get(key int) (*model.Student, error)
	GetList() ([]model.Student, error)
	Edit(key int, student model.Student) (sql.Result, error)
	Save(student model.Student) (sql.Result, error)
}
