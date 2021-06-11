package services

import (
	"github.com/amexws/abcbank/models"
)

func CreateDirectDebit() models.DirectDebit {
	directDebit := models.DirectDebit{
		models.Date{20, 02, 2021},
	}
	return directDebit
}
