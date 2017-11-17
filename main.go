package main

import (
	"github.com/prometheus/common/log"
	"fmt"
	"ajar-go/repository/services"
	"ajar-go/conf"
	"ajar-go/models"
)

func main() {
	db := conf.GetConnection()
	studentService := service.NewStudentService(db)
	// Get Single data
	student, _ := studentService.Get(1)
	fmt.Println("single")
	fmt.Println(student.Id, student.Name)

	students, err := studentService.GetList()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("list")
	for i := range students {
		fmt.Println(students[i].Id, students[i].Name)
	}

	fmt.Println("add")
	studentService.Save(model.Student{1, "Fairuz"})

	fmt.Println("edit")
	studentService.Edit(4, model.Student{1, "Fawwaz"})
}
