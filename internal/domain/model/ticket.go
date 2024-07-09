package model

import "errors"

type Ticket struct {
	ID      int64   `json:"id"`
	Points  Points  `json:"points"`
	Company Company `json:"company"`
	Config  Config  `json:"config"`
}

func (t *Ticket) AddPoints(points int64) (*Ticket, string, error) {
	if points < 0 {
		return nil, "", errors.New("points must be positive")
	}

	if t.Points.CurrentPoints+points > t.Config.MaxPoints {
		remainingPoints := t.Points.CurrentPoints + points - t.Config.MaxPoints
		t.Points.CurrentPoints = t.Config.MaxPoints

		newTicket := &Ticket{
			Points:  Points{CurrentPoints: remainingPoints},
			Company: t.Company,
			Config:  t.Config,
		}

		message := "Max points reached. Prize awarded!"
		return newTicket, message, nil
	}

	t.Points.CurrentPoints += points
	return nil, "", nil
}
