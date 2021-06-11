package services

import (
	"github.com/amexws/abcbank/models"
)

func CreateExternalTransaction() models.ExternalTransaction {
	externalTransaction := models.ExternalTransaction{
		"Indira Nagar Branch",
		"Indira Nagar",
		284678,
		"IFSC80236567",
	}
	return externalTransaction
}
