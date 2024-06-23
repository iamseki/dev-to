package usecases

import (
	"github.com/iamseki/dev-to/apps/mock-tests-in-golang/domain"
	"github.com/iamseki/dev-to/apps/mock-tests-in-golang/usecases/protocols"
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
