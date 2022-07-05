package Models

type Subject struct {
	SubjectID  int
	Code       string
	Name       string
	TeacherIDs []int
}
