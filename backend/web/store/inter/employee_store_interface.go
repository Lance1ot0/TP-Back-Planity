package inter

import (
	"TP-Back-Planity/web/models"
)

type EmployeeStoreInterface interface {
	AddEmployee(employee models.Employee) (int, error)
}
