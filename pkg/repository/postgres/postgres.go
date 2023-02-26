package postgres

import (
	"dockerized/api/pkg/domain"
	"dockerized/api/pkg/repository"
	tools "dockerized/api/tools/postgres"
)

type database struct {
	client *tools.Client
}

func NewRepository(client *tools.Client) repository.Repository {
	return &database{
		client: client,
	}
}

func (r *database) GetRecords() ([]domain.Record, error) {
	rows, err := r.client.DB.Query("SELECT first_name, last_name FROM names")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []domain.Record

	for rows.Next() {
		var rec domain.Record
		if err := rows.Scan(&rec.First_name, &rec.Last_name); err != nil {
			return nil, err
		}
		records = append(records, rec)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return records, nil

}
