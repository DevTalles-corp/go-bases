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

func AddItems(order *Order, items ...Item) {
	order.Items = append(order.Items, items...)
}

func FindItem(order Order, sku string) (Item, bool) {
	for _, item := range order.Items {
		if item.SKU == sku {
			return item, true
		}
	}

	return Item{}, false
}

func GetMeta(order Order, key string) (string, bool) {
	if order.Meta == nil {
		return "", false
	}

	value, ok := order.Meta[key]

	return value, ok
}

func IndexOfItem(order Order, sku string) (int, bool) {
	for index, item := range order.Items {
		if item.SKU == sku {
			return index, true
		}
	}

	return -1, false
} //MS-003 -> 4 true

func Compute(order Order) (t Totals, err error) {
	if err = ValidateOrder(order); err != nil {
		return Totals{}, err
	}

	t.Subtotal = CalcSubtotal(order)
	t.Total = t.Subtotal - t.Discount + t.Tax + t.Shipping

	return t, nil // return

}
