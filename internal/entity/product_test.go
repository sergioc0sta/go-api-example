package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("tone", 10.1)

	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.ID)
	assert.NotEmpty(t, product.CreatedAt)
	assert.Equal(t, "tone", product.Name)
	assert.Equal(t, 10.1, product.Price)
}

func TestNameProductRequeire(t *testing.T) {
	product, err := NewProduct("", 10)
	assert.Nil(t, product)
	assert.Equal(t, ErrNameIsRequire, err)
}

func TestNamePriceRequeire(t *testing.T) {
	product, err := NewProduct("tone", 0)
	assert.Nil(t, product)
	assert.Equal(t, ErrPriceIsRequire, err)
}

func TestNamePriceNotValide(t *testing.T) {
	product, err := NewProduct("tone", -19)
	assert.Nil(t, product)
	assert.Equal(t, ErrPriceIsInvalid, err)
}

func TestProductValidate(t *testing.T) {
	product, err := NewProduct("tone", 12.00)
  println(product.Name)
	assert.NotNil(t, product)
	assert.Nil(t, err)
	assert.Nil(t, product.Validator())
}
