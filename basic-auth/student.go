package main

func init() {
	students = append(students, &Student{"01", "Andromeda", 3})
	students = append(students, &Student{"02", "Budi", 3})
	students = append(students, &Student{"03", "Joko", 3})
}

var students = []*Student{}

type Student struct {
	Id    string
	Name  string
	Grade int
}

func GetStudents() []*Student {
	return students
}

func SelectStudent(id string) *Student {
	for _, student := range students {
		if student.Id == id {
			return student
		}
	}

	return nil
}
