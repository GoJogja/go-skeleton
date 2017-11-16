package main

import (
	"ajar-go/repositories"
	"fmt"
)

func main() {
	student := repositories.GetStudents(2)
	fmt.Println(student.Id, student.Name)
}
