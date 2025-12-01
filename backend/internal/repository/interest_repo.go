package repository

import (
	"LinhuaLink/backend/internal/model"
	"database/sql"
)

type InterestRepository interface {
	GetInterest() ([]model.Interest, error)
}

type interestRepo struct {
	db *sql.DB
}

func NewInterestRepository(db *sql.DB) InterestRepository {
	return &interestRepo{db: db}
}

func (r *interestRepo) GetInterest() ([]model.Interest, error) {
	rows, err := r.db.Query("SELECT * FROM interest")
	if err != nil {
		return []model.Interest{}, err
	}

	var interests []model.Interest
	for rows.Next() {
		var interest model.Interest
		err := rows.Scan(&interest.Name, &interest.Category, &interest.Popularity)
		if err != nil {
			return []model.Interest{}, err
		}
		interests = append(interests, interest)
	}
	return interests, nil
}
