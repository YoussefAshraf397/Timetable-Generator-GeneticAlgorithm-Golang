package Models

type Session struct {
	SessionID   int `json:"session_id" gorm:"primary_key"`
	ClassroomID int `json:"classroom_id"`
	SubjectID   int `json:"subject_id"`
	TeacherID   int `json:"teacher_id"`
	TimeSlotID  int `json:"time_slot_id"`
}
