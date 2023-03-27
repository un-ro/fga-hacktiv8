package models

type Student struct {
	Id    string
	Name  string
	Grade int32
}

var students = []*Student{}

func GetStudents() []*Student {
	return students
}

func SelectStudent(id string) *Student {
	for _, s := range students {
		if s.Id == id {
			return s
		}
	}
	return nil
}

func init() {
	students = append(students, &Student{"s001", "Ethan", 2})
	students = append(students, &Student{"s002", "Bella", 3})
	students = append(students, &Student{"s003", "James", 2})
}
