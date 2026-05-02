package events

import (
	"context"

	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepository {
	return &EventRepository{db: db}
}

func (repo *EventRepository) Create(ctx context.Context, event *Event) error {
	return gorm.G[Event](repo.db).Create(ctx, event)
}

func (repo *EventRepository) GetById(ctx context.Context, id uint) (Event, error) {
	return gorm.G[Event](repo.db).Where("id = ?", id).First(ctx)
}

func (repo *EventRepository) Update(ctx context.Context, id uint, updates map[string]any) error {
	_, err := gorm.G[map[string]any](repo.db).Where("id = ?", id).Updates(ctx, updates)
	return err
}

func (repo *EventRepository) Delete(ctx context.Context, id uint) error {
	_, err := gorm.G[Event](repo.db).Where("id = ?", id).Delete(ctx)
	return err
}
