package models

import (
	"context"
	"time"
)

type Ticket struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	EventID   uint      `json:"eventId"`
	Event     Event     `gorm:"foreignKey:EventID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"event"`
	Entered   bool      `json:"entered" default:"false"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TicketRepository interface {
	GetMany(ctx context.Context) ([]*Ticket, error)
	GetOne(ctx context.Context, ticketId uint) (*Ticket, error)
	CreateOne(ctx context.Context, ticket *Ticket) (*Ticket, error)
	UpdateOne(ctx context.Context, ticketId uint, updateData map[string]interface{}) (*Ticket, error)
}

type ValidateTicket struct {
	TicketID uint `json:"ticketId"`
}
