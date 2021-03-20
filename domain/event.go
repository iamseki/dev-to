package domain

import "time"

type Event struct {
	Title    string
	Date     time.Time
	Place    string
	Category string
	KeyWords []string
}

type Filter struct {
	Title string
}

type EventSaver interface {
	Save(Event) error
}

type EventFinder interface {
	Find(Filter) ([]Event, error)
}
