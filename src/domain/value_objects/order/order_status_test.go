package domain_value_object

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateOrderStatus(t *testing.T) {
	result := CreateOrderStatus(uint8(3))

	assert.NotEqual(t, OrderStatus(StatusDraft), result, "result should not be equals to draft")
	assert.NotEqual(t, OrderStatus(StatusDone), result, "result should not be equals to done")
}

func TestOrderStatusToString(t *testing.T) {
	result := OrderStatus(StatusDraft).ToString()

	assert.Equal(t, "Draft", result, "result should be equals to 'Draft'")
}

func TestOrderStatusEnumIndex(t *testing.T) {
	result := OrderStatus(StatusDraft).EnumIndex()

	assert.Equal(t, uint8(1), result, "result should be equals to 1")
}

func TestIsValid_True(t *testing.T) {
	result := OrderStatus(StatusDraft).IsValid()

	assert.True(t, result)
}

func TestIsValid_False(t *testing.T) {
	result := OrderStatus(0).IsValid()

	assert.False(t, result)
}
