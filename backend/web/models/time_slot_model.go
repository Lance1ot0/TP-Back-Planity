package models

type TimeSlot struct {
	TimeSlotID     int    `json:"timeSlotID"`
	Date           string `json:"date"`
	StartTime      string `json:"startTime"`
	EndTime        string `json:"endTime"`
	ProfessionalID int    `json:"professionalID"`
	EmployeeID     int    `json:"employeeID"`
	AvailabilityID int    `json:"availabilityID"`
}
