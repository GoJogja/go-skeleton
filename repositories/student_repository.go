package repositories

import (
	"ajar-go/conf"
	"ajar-go/models"
	"github.com/prometheus/common/log"
	"strconv"
)

var db = conf.GetConnection()

func GetStudents(key int) *models.Student {
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
		return &models.Student{id, name}
	}
	return &models.Student{0, ""}
}
