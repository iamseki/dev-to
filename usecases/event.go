package usecases

import (
	"github.com/iamseki/dev-to/domain"
	"github.com/iamseki/dev-to/usecases/protocols"
)

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
