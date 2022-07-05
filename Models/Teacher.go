package Models

type Teacher struct {
	TeacherID   int    `json:"teacher_id" gorm:"primary_key"`
	TeacherName string `json:"teacher_name"`
}
