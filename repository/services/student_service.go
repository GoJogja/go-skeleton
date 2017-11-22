package service

import (
	"ajar-go/models"
	"github.com/prometheus/common/log"
	"strconv"
	"database/sql"
)

type StudentRepository interface {
	Get(key int) (*model.Student, error)
	GetList() ([]model.Student, error)
	Edit(key int, student model.Student) (sql.Result, error)
	Save(student model.Student) (sql.Result, error)
	Delete(key int) (sql.Result, error)
}

type studentRepository struct {
	db *sql.DB
}

func NewStudentService(db *sql.DB) StudentRepository {
	return &studentRepository{db}
}

func (conn *studentRepository) Get(key int) (*model.Student, error) {
	row, err := conn.db.Query("select id, name from students where id = " + strconv.Itoa(key))
	if err != nil {
		log.Fatal(err)
	}

	var (
		id   int
		name string
	)
	for row.Next() {
		row.Scan(&id, &name)
	}
	return &model.Student{id, name}, nil
}

func (conn *studentRepository) GetList() ([]model.Student, error) {
	var students []model.Student

	row, err := conn.db.Query("select id, name from students")
	if err != nil {
		log.Fatal(err)
	}

	st := model.Student{}

	for row.Next() {
		row.Scan(&st.Id, &st.Name)
		students = append(students, st)
	}
	return students, nil
}

func (conn *studentRepository) Edit(key int, student model.Student) (sql.Result, error) {
	stmt, err := conn.db.Prepare("update students set name = ? where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(student.Name, key)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

func (conn *studentRepository) Save(student model.Student) (sql.Result, error) {
	stmt, err := conn.db.Prepare("insert into students (name) values (?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(student.Name)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

func (conn *studentRepository) Delete(key int) (sql.Result, error) {
	stmt, err := conn.db.Prepare("delete from students where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(key)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}
