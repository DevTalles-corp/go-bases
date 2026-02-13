package checkout

type ShippingFn func(Order) Money

func FreeShipping(Order) Money {
	return 0
}

func NewShippingByZone(zone string) ShippingFn {
	switch zone {
	case "LOCAL":
		return FreeShipping
	case "NATIONAL":
		return func(Order) Money { return 499 }
	default:
		return func(Order) Money { return 2499 }
	}
}
