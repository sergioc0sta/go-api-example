package entity

import (
	"errors"
	"goexpert-api/pkg/entity"
	"time"
)

var (
	ErrIDIsRequire    = errors.New("id is require")
	ErrIDIsInvalid    = errors.New("id is invalid")
	ErrNameIsRequire  = errors.New("name is invalid")
	ErrPriceIsRequire = errors.New("price is invalid")
	ErrPriceIsInvalid = errors.New("price is invalid")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func (p *Product) Validator() error {
	if p.ID.String() == "" {
		return ErrIDIsRequire
	}

	if _, err := entity.ParseUid(p.ID.String()); err != nil {
		return ErrIDIsInvalid
	}

	if p.Name == "" {
		return ErrNameIsRequire
	}

	if p.Price == 0 {
		return ErrPriceIsRequire
	}

	if p.Price < 0 {
		return ErrPriceIsInvalid
	}

	return nil
}

func NewProduct(name string, price float64) (*Product, error) {
	product := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}

	if err := product.Validator(); err != nil {
		return nil, err
	}

	return product, nil

}
