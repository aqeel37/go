package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewItem(t *testing.T) {
	item, err := NewItem(12, "11000110022", "Test Chocolate", "2030-12-31", "10000")
	assert.NoError(t, err)
	assert.Equal(t, int32(12), item.Quantity())
	assert.Equal(t, "11000110022", item.itemNumber)
	assert.Equal(t, "Test Chocolate", item.description)
	assert.Equal(t, "2030-12-31", item.expiryDate)
	assert.Equal(t, "10000", item.lotNumber)
}
