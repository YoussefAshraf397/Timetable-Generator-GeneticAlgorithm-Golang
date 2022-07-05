package Models

type Classroom struct {
	ClassroomID       int   `json:"classroom_id" gorm:"primary_key"`
	ClassroomCapacity int   `json:"classroom_capacity"`
	SubjectIds        []int `json:"subject_ids" gorm:"type:text"`
}

func (g *Classroom) GetSubjectIDs() []int {

	return g.SubjectIds
}

func (g *Classroom) GetClassroomID() int {

	return g.ClassroomID
}
