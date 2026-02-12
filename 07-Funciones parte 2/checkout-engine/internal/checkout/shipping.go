package checkout

type ShippingFn func(Order) Money

func FreeShipping(Order) Money {
	return 0
}
