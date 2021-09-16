package gopher

import (
	"errors"
	"sync"
)

type InMemoryRepository struct {
	gophers []Gopher
	sync.Mutex
}

func NewMemoryRepository() *InMemoryRepository {
	gophers := []Gopher{
		{
			ID:         "1",
			Name:       "Original Gopher",
			Hired:      true,
			Profession: "Logo",
		}, {
			ID:         "2",
			Name:       "Jan",
			Hired:      true,
			Profession: "The Janitor",
		},
	}

	return &InMemoryRepository{
		gophers: gophers,
	}
}

// GetGophers returns all gophers
func (imr *InMemoryRepository) GetGophers() ([]Gopher, error) {
	return imr.gophers, nil
}

// GetGopher will return a goper by its ID
func (imr *InMemoryRepository) GetGopher(id string) (Gopher, error) {
	for _, gopher := range imr.gophers {
		if gopher.ID == id {
			return gopher, nil
		}
	}
	return Gopher{}, errors.New("no such gopher exists")
}
