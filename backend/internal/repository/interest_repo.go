package repository

import (
	"LinhuaLink/backend/internal/model"
	"database/sql"
)

type InterestRepository interface {
	GetInterest() (map[string][]model.Interest, error)
}

type interestRepo struct {
	db *sql.DB
}

func NewInterestRepository(db *sql.DB) InterestRepository {
	return &interestRepo{db: db}
}

func (r *interestRepo) GetInterest() (map[string][]model.Interest, error) {
	rows, err := r.db.Query("SELECT name, category, popularity FROM interest")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	interests := make(map[string][]model.Interest)

	for rows.Next() {
		var res model.Interest
		var category string
		if err := rows.Scan(&res.Name, &category, &res.Popularity); err != nil {
			return nil, err
		}

		interests[category] = append(interests[category], res)
	}

	return interests, nil
}
