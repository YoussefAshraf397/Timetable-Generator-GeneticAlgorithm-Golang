package Models

import "math/rand"

type Subject struct {
	SubjectID  int
	Code       string
	Name       string
	TeacherIDs []int
}

func (m *Subject) GetRandomTeacherID() int {
	var teacherID = m.TeacherIDs[rand.Intn(9494268)%len(m.TeacherIDs)]
	return teacherID
}
