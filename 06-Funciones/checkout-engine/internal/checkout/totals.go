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
