package protocols

import "github.com/iamseki/dev-to/apps/mock-tests-in-golang/domain"

type AddInMemoryRepository interface {
	Add(domain.Event) error
}
