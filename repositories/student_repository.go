package repositories

import (
	"ajar-go/conf"
	"ajar-go/models"
	"github.com/prometheus/common/log"
	"strconv"
)

var db = conf.GetConnection()

func GetStudent(key int) (*models.Student, error) {
	row, err := db.Query("select id, name from students where id = " + strconv.Itoa(key))
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
	return &models.Student{id, name}, nil
}

func GetStudents() ([]models.Student, error) {
	var students []models.Student

	row, err := db.Query("select id, name from students")
	if err != nil {
		log.Fatal(err)
	}

	st := models.Student{}

	for row.Next() {
		row.Scan(&st.Id, &st.Name)
		students = append(students, st)
	}
	return students, nil
}
