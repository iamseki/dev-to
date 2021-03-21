package repository

import (
	"sync"

	"github.com/iamseki/dev-to/domain"
)

type Memo struct {
	storage []domain.Event
}

var singleton *Memo
var once sync.Once

func NewInMemoryRepository() *Memo {
	once.Do(func() {
		singleton = &Memo{
			storage: []domain.Event{
				{Title: "fake title01"},
				{Title: "fake title01"},
				{Title: "fake title"},
				{Title: "fake title"},
			},
		}
	})
	return singleton
}

func (m *Memo) Add(e domain.Event) error {
	m.storage = append(m.storage, e)
	return nil
}

func (m *Memo) Get(title string) ([]domain.Event, error) {
	var res []domain.Event

	if title == "" {
		return m.storage, nil
	}
	for _, e := range m.storage {
		if e.Title == title {
			res = append(res, e)
		}
	}
	return res, nil
}
