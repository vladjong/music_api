package entity

type Song struct {
	Id       int64   `json:"id" db:"id" goqu:"skipinsert"`
	Name     string  `json:"name" db:"name"`
	Duration float64 `json:"duration" db:"duration"`
}
