package entity

type Song struct {
	Id       int64   `json:"id" db:"-" goqu:"skipinsert"`
	Name     string  `json:"name" db:"name"`
	Duration float64 `json:"duration" db:"duration"`
}
