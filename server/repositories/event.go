package repositories

import (
	"context"
	"time"

	"github.com/Dunsin-cyber/ticbuk/models"
)

type EventRepository struct {
	db any
}

func NewEventRepository(db any) models.EventRepository {
	return &EventRepository{
		db: db,
	}
}

func (er *EventRepository) GetMany(ctx context.Context) ([]*models.Event, error) {
	events := []*models.Event{}
	events = append(events, &models.Event{
		ID:        1,
		Name:      "Sample Event",
		Location:  "Sample Location",
		Date:      time.Now(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	return events, nil
}

func (er *EventRepository) GetOne(ctx context.Context, eventId string) (*models.Event, error) {
	return nil, nil
}

func (er *EventRepository) CreateOne(ctx context.Context, event models.Event) (*models.Event, error) {
	return nil, nil
}
