package domain_value_object

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeOrderItem() OrderItem {
	return CreateOrderItem(
		"SGM ANANDA",
		uint8(10),
		float32(100000),
	)
}

func TestCreateOrderItem(t *testing.T) {
	result := CreateOrderItem(
		"SGM ANANDA",
		uint8(10),
		float32(100000),
	)

	assert.Equal(t, "SGM ANANDA", result.name, "name should be equal")
	assert.Equal(t, uint8(10), result.qty, "qty should be equal")
	assert.Equal(t, float32(100000), result.price, "price should be equal")
	assert.Equal(t, float32(1_000_000), result.total, "total should be equal from qty * price")
}

func TestOrderItemEqualTo(t *testing.T) {
	current := makeOrderItem()
	other := makeOrderItem()
	result := current.EqualTo(other)

	assert.Equal(t, true, result, "order item should be equal")
}

func TestOrderItemEqualToInvalid(t *testing.T) {
	current := makeOrderItem()
	other := CreateOrderItem(
		"SGM",
		uint8(100),
		float32(1_000_000),
	)
	result := current.EqualTo(other)

	assert.Equal(t, false, result, "order item should be different")
}

func TestGetOrderItemName(t *testing.T) {
	entity := makeOrderItem()
	result := entity.GetName()

	assert.Equal(t, "SGM ANANDA", result, "name should be equal")
}

func TestGetOrderItemQty(t *testing.T) {
	entity := makeOrderItem()
	result := entity.GetQty()

	assert.Equal(t, uint8(10), result, "qty should be equal")
}

func TestGetOrderItemPrice(t *testing.T) {
	entity := makeOrderItem()
	result := entity.GetPrice()

	assert.Equal(t, float32(100_000), result, "price should be equal")
}

func TestGetOrderItemTotalPrice(t *testing.T) {
	entity := makeOrderItem()
	result := entity.GetTotalPrice()

	assert.Equal(t, float32(1_000_000), result, "totalPrice should be equal")
}

func TestOrderItemMarshalJSON(t *testing.T) {
	entity := makeOrderItem()
	result, err := entity.MarshalJSON()

	if err != nil {
		assert.Fail(t, "error should be nil")
	} else {
		assert.Equal(t, true, len(result) > 0, "marshal result should be have length more than 0")
	}
}
