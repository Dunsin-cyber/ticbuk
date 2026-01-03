package models

import (
	"context"
	"time"
)

type Ticket struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	EventID   uint      `json:"eventId"`
	Event     Event     `gorm:"foreignKey:EventID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"event"`
	UserID    uint      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"userId"`
	Entered   bool      `json:"entered" default:"false"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TicketRepository interface {
	GetMany(ctx context.Context, userId uint) ([]*Ticket, error)
	GetOne(ctx context.Context,  userId uint, ticketId uint) (*Ticket, error)
	CreateOne(ctx context.Context,  userId uint, ticket *Ticket) (*Ticket, error)
	UpdateOne(ctx context.Context,  userId uint, ticketId uint, updateData map[string]interface{}) (*Ticket, error)
}

type ValidateTicket struct {
	TicketID uint `json:"ticketId"`
	OwnerID  uint `json:"ownerId"`
}
