package model

type Config struct {
	ID        int64   `json:"id"`
	Prizes    []Prize `json:"prizes"`
	MaxPoints int64   `json:"max_points"`
}
