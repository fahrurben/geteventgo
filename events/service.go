package events

import "context"

type EventService struct {
	repo *EventRepository
}

func (s *EventService) Create(ctx context.Context, model *Event) error {
	return s.repo.Create(ctx, model)
}
