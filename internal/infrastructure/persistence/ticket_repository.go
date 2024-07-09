package persistence

import (
	"database/sql"
	"errors"

	"github.com/ananiaslitz/go-api-ticket/internal/domain/model"
	"github.com/ananiaslitz/go-api-ticket/internal/infrastructure/repository"
)

type TicketRepositoryImpl struct {
	DB *sql.DB
}

func NewTicketRepository(db *sql.DB) repository.TicketRepository {
	return &TicketRepositoryImpl{DB: db}
}

func (r *TicketRepositoryImpl) Create(ticket *model.Ticket) error {
	query := `INSERT INTO tickets (points, company_id, config_id) VALUES ($1, $2, $3) RETURNING id`
	err := r.DB.QueryRow(query, ticket.Points.CurrentPoints, ticket.Company.ID, ticket.Config.ID).Scan(&ticket.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *TicketRepositoryImpl) GetByID(id int64) (*model.Ticket, error) {
	var ticket model.Ticket
	query := `SELECT id, points, company_id, config_id FROM tickets WHERE id = $1`
	row := r.DB.QueryRow(query, id)
	err := row.Scan(&ticket.ID, &ticket.Points.CurrentPoints, &ticket.Company.ID, &ticket.Config.ID)
	if err == sql.ErrNoRows {
		return nil, errors.New("ticket not found")
	} else if err != nil {
		return nil, err
	}
	return &ticket, nil
}

func (r *TicketRepositoryImpl) GetAll() ([]*model.Ticket, error) {
	query := `SELECT id, points, company_id, config_id FROM tickets`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tickets []*model.Ticket
	for rows.Next() {
		var ticket model.Ticket
		err := rows.Scan(&ticket.ID, &ticket.Points.CurrentPoints, &ticket.Company.ID, &ticket.Config.ID)
		if err != nil {
			return nil, err
		}
		tickets = append(tickets, &ticket)
	}

	return tickets, nil
}

func (r *TicketRepositoryImpl) Update(ticket *model.Ticket) error {
	query := `UPDATE tickets SET points = $1 WHERE id = $2`
	_, err := r.DB.Exec(query, ticket.Points.CurrentPoints, ticket.ID)
	return err
}

func (r *TicketRepositoryImpl) Delete(id int64) error {
	query := `DELETE FROM tickets WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	return err
}
