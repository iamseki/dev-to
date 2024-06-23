package usecases_test

import (
	"testing"

	"github.com/iamseki/dev-to/apps/mock-tests-in-golang/domain"
	"github.com/iamseki/dev-to/apps/mock-tests-in-golang/usecases"
)

type addEventFakeRepository struct {
	MockAddFn func(domain.Event) error
}

func (fake *addEventFakeRepository) Add(e domain.Event) error {
	return fake.MockAddFn(e)
}

func newAddEventFakeRepository() *addEventFakeRepository {
	return &addEventFakeRepository{
		MockAddFn: func(e domain.Event) error { return nil },
	}
}

func TestAddEventInMemorySucceed(t *testing.T) {
	r := newAddEventFakeRepository()
	sut := usecases.NewAddEventInMemory(r)

	err := sut.Save(domain.Event{})
	if err != nil {
		t.Error("Expect error to be nil but got:", err)
	}
}
