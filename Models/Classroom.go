package Models

type Classroom struct {
	ClassroomID       int
	ClassroomCapacity int
	SubjectIds        []int
}

func (g *Classroom) GetSubjectIDs() []int {

	return g.SubjectIds
}

func (g *Classroom) GetClassroomID() int {

	return g.ClassroomID
}
