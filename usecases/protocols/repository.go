package protocols

import "github.com/iamseki/dev-to/domain"

type AddInMemoryRepository interface {
	Add(domain.Event) error
}

type FindInMemoryRepository interface {
	Find(domain.Filter) ([]domain.Event, error)
}
