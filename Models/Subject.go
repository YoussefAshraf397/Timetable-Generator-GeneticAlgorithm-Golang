package Models

import (
	"math/rand"
)

type Subject struct {
	SubjectID  int    `json:"subject_id" gorm:"primary_key"`
	Code       string `json:"code"`
	Name       string `json:"name"`
	TeacherIDs []int  `json:"teacher_ids" gorm:"type:text"`
}

func (m *Subject) GetRandomTeacherID() int {
	var teacherID = m.TeacherIDs[rand.Intn(9494268)%len(m.TeacherIDs)]
	return teacherID
}
