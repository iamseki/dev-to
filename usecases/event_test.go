package usecases_test

import (
	"errors"
	"testing"

	"github.com/iamseki/dev-to/domain"
	"github.com/iamseki/dev-to/usecases"
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

type findEventFakeRepository struct {
	MockFindFn func(string) ([]domain.Event, error)
}

func (fake *findEventFakeRepository) Get(title string) ([]domain.Event, error) {
	return fake.MockFindFn(title)
}

func newFindEventFakeRepository() *findEventFakeRepository {
	return &findEventFakeRepository{
		MockFindFn: func(title string) ([]domain.Event, error) {
			e := []domain.Event{{Title: "fake event0"}, {Title: "fake event01"}}
			return e, nil
		},
	}
}

func TestFindEventInMemorySucceed(t *testing.T) {
	r := newFindEventFakeRepository()
	u := usecases.NewFindEventInMemory(r)

	events, err := u.Find(domain.Filter{Title: "fake search"})
	if err != nil {
		t.Error("Expect err to be nil but got:", err)
	}

	if len(events) != 2 {
		t.Error("Expect len(events) equals to 2 but got:", len(events))
	}
}

func TestFindEventInMemoryFailure(t *testing.T) {
	r := newFindEventFakeRepository()
	r.MockFindFn = func(title string) ([]domain.Event, error) { return []domain.Event{}, errors.New("Fake Error") }
	u := usecases.NewFindEventInMemory(r)

	events, err := u.Find(domain.Filter{Title: "fake search"})
	if len(events) > 0 {
		t.Error("Expect len(events) equals to 0 but got:", events)
	}
	if err == nil {
		t.Error("Expect err != nil but got:", err)
	}
}
