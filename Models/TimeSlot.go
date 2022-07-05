package Models

type TimeSlot struct {
	TimeSlotID int    `json:"time_slot_id" gorm:"primary_key"`
	TimeSlot   string `json:"time_slot"`
}
