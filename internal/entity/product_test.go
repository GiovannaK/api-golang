package entity

import (
	"github.com/stretchr/testify/assert" 
	"testing"
)

func TestNewProduc(t *testing.T) {
	p, err := NewProduct("Product 1", 10.00)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.ID)
	assert.Equal(t, "Product 1", p.Name)
	assert.Equal(t, 10.00, p.Price)

}

func TestProductWhenIsRequired(t *testing.T) {
	p, err := NewProduct("", 10)
	assert.Nil(t, p)
	assert.Equal(t, ErrNameRequired, err)

}

func TestProductWhenPriceIsZero(t *testing.T) {
	p, err := NewProduct("Product 1", 0)
	assert.Nil(t, p)
	assert.Equal(t, ErrPriceRequired, err)
}

func TestProductWhenPriceIsNegative(t *testing.T) {
	p, err := NewProduct("Product 1", -10)
	assert.Nil(t, p)
	assert.Equal(t, ErrInvalidPrice, err)
}
