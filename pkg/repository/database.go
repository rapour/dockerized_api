package repository

import (
	"dockerized/api/pkg/domain"
)

type Repository interface {
	GetRecords() ([]domain.Record, error)
}
