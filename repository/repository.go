package repository

import (
    "github.com/desmondrawls/rock-paper-scissors/history"
    "github.com/desmondrawls/rock-paper-scissors/models"
    "github.com/desmondrawls/rock-paper-scissors/play"
)

type Repository interface {
    play.Repository
    history.Repository
}

type FakeRepository struct {
    records []models.Record
}

func (r *FakeRepository) List() ([]models.Record, error) {
    return r.records, nil
}

func (r *FakeRepository) Save(record models.Record) error {
    r.records = append(r.records, record)
    return nil
}
