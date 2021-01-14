package backend

import (
	"fmt"
)

type Category string

const (
	FMCG    Category = "fmcg"
	Textile Category = "textile"
)

type Product struct {
	ID       string
	Name     string  `json:"name"`
	Brand    string  `json:"brand"`
	Quantity int     `json:"qty"`
	Price    float64 `json:"price"`
	Category `json:"category"`
}

func (c Category) ToString() string {
	return string(c)
}

func (c Category) Validate() error {
	switch c {
	case FMCG:
	case Textile:
		return nil
	default:
		return fmt.Errorf("Unsupported category")
	}

	return nil
}

func (p Product) Validate() error {
	if p.ID == "" {
		return fmt.Errorf("invalid id")
	}

	if p.Price < 0 {
		return fmt.Errorf("Invalid price")
	}

	if p.Quantity < 0 {
		return fmt.Errorf("Invalid quantity")
	}

	return p.Category.Validate()
}
