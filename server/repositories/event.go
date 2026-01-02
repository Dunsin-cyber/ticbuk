package repositories

import (
	"context"

	"github.com/Dunsin-cyber/ticbuk/models"
	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) models.EventRepository {
	return &EventRepository{
		db: db,
	}
}

func (er *EventRepository) GetMany(ctx context.Context) ([]*models.Event, error) {
	events := []*models.Event{}

	res := er.db.Model(&models.Event{}).WithContext(ctx).Find(&events)

	if res.Error != nil {
		return nil, res.Error
	}
	return events, nil
}

func (er *EventRepository) GetOne(ctx context.Context, eventId string) (*models.Event, error) {
	return nil, nil
}

func (er *EventRepository) CreateOne(ctx context.Context, event models.Event) (*models.Event, error) {
	return nil, nil
}
