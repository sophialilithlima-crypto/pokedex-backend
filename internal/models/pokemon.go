package models

type Pokemon struct {
	ID uint `gorm:"primaryKey"`

	PokedexID int

	Name string

	Type1 string

	Type2 string

	Height float64

	Weight float64

	Description string
}