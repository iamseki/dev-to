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
		singleton = &Memo{}
	})
	return singleton
}

func (m *Memo) Add(e domain.Event) error {
	m.storage = append(m.storage, e)
	return nil
}

func (m *Memo) Get(title string) []domain.Event {
	var res []domain.Event
	for _, e := range m.storage {
		if e.Title == title {
			res = append(res, e)
		}
	}
	return res
}
