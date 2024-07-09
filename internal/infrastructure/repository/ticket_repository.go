package repository

import "github.com/ananiaslitz/go-api-ticket/internal/domain/model"

type TicketRepository interface {
	Create(ticket *model.Ticket) error
	GetByID(id int64) (*model.Ticket, error)
	GetAll() ([]*model.Ticket, error)
	Update(ticket *model.Ticket) error
	Delete(id int64) error
}
