package domain_value_object

import "encoding/json"

type OrderItem struct {
	name  string
	qty   uint8
	price float32
	total float32
}

func CreateOrderItem(name string, qty uint8, price float32) OrderItem {
	return OrderItem{
		name:  name,
		qty:   qty,
		price: price,
		total: float32(qty) * price,
	}
}

func (o OrderItem) EqualTo(other OrderItem) bool {
	return o.GetName() == other.GetName() && o.GetQty() == other.GetQty() &&
		o.GetPrice() == other.GetPrice() && o.GetTotalPrice() == other.GetTotalPrice()
}

func (o OrderItem) GetName() string {
	return o.name
}

func (o OrderItem) GetQty() uint8 {
	return o.qty
}

func (o OrderItem) GetPrice() float32 {
	return o.price
}

func (o OrderItem) GetTotalPrice() float32 {
	return o.total
}

func (o OrderItem) MarshalJSON() ([]byte, error) {
	j, _ := json.Marshal(struct {
		Name  string
		Qty   uint8
		Price float32
		Total float32
	}{
		Name:  o.name,
		Qty:   o.qty,
		Price: o.price,
		Total: o.total,
	})

	return j, nil
}
