package checkout

func NewOrder(id, customer string) Order {
	return Order{
		ID:       id,
		Customer: customer,
		Items:    []Item{},
		Meta:     map[string]string{},
	}
}

func AddItem(o *Order, item Item) {
	o.Items = append(o.Items, item)
}

func RemoveItem(o *Order, sku string) bool {
	for i := range o.Items {
		if o.Items[i].SKU == sku { // "b" = "b"
			o.Items = append(o.Items[:i], o.Items[i+1:]...) // ["a", "b", "c", "d"]
			return true
		}
	}

	return false
}

func CalcLineTotal(item Item) Money {
	return item.Price * Money(item.Qty)
}

func CalcSubtotal(order Order) Money {
	var sum Money
	for _, item := range order.Items {
		sum += CalcLineTotal(item) // sum = sum + CalcLineTotal(item)
	}
	return sum
}

func CalcTotalQty(order Order) int {
	total := 0
	for _, item := range order.Items {
		total += item.Qty
	}
	return total
}
