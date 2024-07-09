package model

type Prize struct {
	Name           string `json:"name"`
	RequiredPoints int64  `json:"required_points"`
}
