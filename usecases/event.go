package usecases

import (
	"github.com/iamseki/dev-to/domain"
	"github.com/iamseki/dev-to/usecases/protocols"
)

// AddEventInMemory resources
type AddEventInMemory struct {
	repository protocols.AddInMemoryRepository
}

func (usecase *AddEventInMemory) Save(e domain.Event) error {
	err := usecase.repository.Add(e)
	return err
}

func NewAddEventInMemory(r protocols.AddInMemoryRepository) *AddEventInMemory {
	return &AddEventInMemory{repository: r}
}

// FindEventInMemory resources

type FindEventInMemory struct {
	repository protocols.FindInMemoryRepository
}

func (usecase *FindEventInMemory) Find(f domain.Filter) ([]domain.Event, error) {
	events, err := usecase.repository.Get(f.Title)
	if err != nil {
		return []domain.Event{}, err
	}
	return events, nil
}

func NewFindEventInMemory(r protocols.FindInMemoryRepository) *FindEventInMemory {
	return &FindEventInMemory{repository: r}
}
