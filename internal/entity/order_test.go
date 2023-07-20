package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfGetAnErrorIfIdIsBlank(t *testing.T) {
	order := Order{}
	assert.Error(t, order.Validate(), "id is required")
}

func TestIfGetAnErrorWhenPriceIsBlank(t *testing.T) {
	order := Order{ID: "123"}
	assert.Error(t, order.Validate(), "price must be grater than zero")
}

func TestIfGetAnErrorWhenPriceIsZero(t *testing.T) {
	order := Order{ID: "123", Price: 0}
	assert.Error(t, order.Validate(), "price must be grater than zero")
}

func TestIfGetAnErrorWhenTaxIsBlank(t *testing.T) {
	order := Order{ID: "123", Price: 10}
	assert.Error(t, order.Validate(), "tax must be grater than zero")
}

func TestFinalPrice(t *testing.T) {
	order := Order{ID: "123", Price: 10, Tax: 1}
	assert.NoError(t, order.Validate())
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 1.0, order.Tax)
	order.CalculateFinalPrice()
	assert.Equal(t, 11.0, order.FinalPrice)
}
