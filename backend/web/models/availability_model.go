package models

type Availability struct {
	AvailabilityID int    `json:"availabilityID"`
	EmployeeID     int    `json:"employeeID"`
	DayOfWeek      string `json:"dayOfWeek"`
	StartTime      string `json:"startTime"`
	EndTime        string `json:"endTime"`
}
