package models

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Event struct {
	ID                    uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name                  string    `json:"name"`
	Location              string    `json:"location"`
	TotalTicketsPurchased int64     `json:"totalTicketsPurchased" gorm:"-"`
	TotalTicketsEntered   int64     `json:"totalTicketsEntered" gorm:"-"`
	Date                  time.Time `json:"date"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}

type EventRepository interface {
	GetMany(ctx context.Context) ([]*Event, error)
	GetOne(ctx context.Context, eventId uint) (*Event, error)
	CreateOne(ctx context.Context, event *Event) (*Event, error)
	UpdateOne(ctx context.Context, eventId uint, updateData map[string]interface{}) (*Event, error)
	DeleteOne(ctx context.Context, eventId uint) error
}

func (e *Event) AfterFind(db *gorm.DB) (err error) {
	baseQuery := db.Model(&Ticket{}).Where("event_id = ?", e.ID)

	if res := baseQuery.Count(&e.TotalTicketsPurchased); res.Error != nil {
		return res.Error
	}

	if res := baseQuery.Where("entered = ?", true).Count(&e.TotalTicketsEntered); res.Error != nil {
		return res.Error
	}

	return nil
}
