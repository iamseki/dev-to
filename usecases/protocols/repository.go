package protocols

import "github.com/iamseki/dev-to/domain"

type AddInMemoryRepository interface {
	Add(domain.Event) error
}

type FindInMemoryRepository interface {
	Get(string) ([]domain.Event, error)
}
