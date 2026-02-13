package checkout

type TaxFn func(Order) Money

func NoTax(Order) Money {
	return 0
}

func IVA16(order Order) Money {
	sub := order.CalcSubtotal()
	return sub * 16 / 100
}

func NewTaxByState(state string) TaxFn {
	switch state {
	case "CDMX":
		return func(o Order) Money { return o.CalcSubtotal() * 16 / 100 }
	case "NL":
		return func(o Order) Money { return o.CalcSubtotal() * 15 / 100 }
	case "QRO":
		return func(o Order) Money { return o.CalcSubtotal() * 20 / 100 }
	case "GDL":
		return func(o Order) Money { return o.CalcSubtotal() * 14 / 100 }
	default:
		return NoTax
	}
}
