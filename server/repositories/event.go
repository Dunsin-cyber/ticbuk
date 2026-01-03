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

	res := er.db.Model(&models.Event{}).WithContext(ctx).Order("updated_at desc").Find(&events)

	if res.Error != nil {
		return nil, res.Error
	}
	return events, nil
}

func (er *EventRepository) GetOne(ctx context.Context, eventId uint) (*models.Event, error) {
	event := &models.Event{}

	res := er.db.Model(event).WithContext(ctx).Where("id = ?", eventId).First(event)

	if res.Error != nil {
		return nil, res.Error
	}
	return event, nil
}

func (er *EventRepository) CreateOne(ctx context.Context, event *models.Event) (*models.Event, error) {
	res := er.db.Model(event).WithContext(ctx).Create(event)

	if res.Error != nil {
		return nil, res.Error
	}
	return event, nil
}

func (er *EventRepository) UpdateOne(ctx context.Context, eventId uint, updateData map[string]interface{}) (*models.Event, error) {
	event := &models.Event{}

	updateRes := er.db.Model(event).Where("id = ?", eventId).Updates(updateData)

	if updateRes.Error != nil {
		return nil, updateRes.Error
	}

	getRes := er.db.Model(event).Where("id = ?", eventId).First(event)

	if getRes.Error != nil {
		return nil, getRes.Error
	}
	return event, nil
}

func (er *EventRepository) DeleteOne(ctx context.Context, eventId uint) error {
	res := er.db.Model(&models.Event{}).WithContext(ctx).Delete(&models.Event{}, eventId)
	return res.Error
}
