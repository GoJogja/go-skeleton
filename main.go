package main

import (
	"ajar-go/repositories"
	"github.com/prometheus/common/log"
	"fmt"
)

func main() {
	student, _ := repositories.GetStudent(1)
	fmt.Println(student.Id, student.Name)

	students, err := repositories.GetStudents()
	if err != nil {
		log.Fatal(err)
	}

	for i := range students {
		fmt.Println(students[i].Id, students[i].Name)
	}
}
