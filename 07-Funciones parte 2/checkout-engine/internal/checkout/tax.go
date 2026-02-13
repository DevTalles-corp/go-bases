package checkout

type TaxFn func(Order) Money

func NoTax(Order) Money {
	return 0
}

func IVA16(order Order) Money {
	sub := CalcSubtotal(order)
	return sub * 16 / 100
}

func NewTaxByState(state string) TaxFn {
	switch state {
	case "CDMX":
		return func(o Order) Money { return CalcSubtotal(o) * 16 / 100 }
	case "NL":
		return func(o Order) Money { return CalcSubtotal(o) * 15 / 100 }
	case "QRO":
		return func(o Order) Money { return CalcSubtotal(o) * 20 / 100 }
	case "GDL":
		return func(o Order) Money { return CalcSubtotal(o) * 14 / 100 }
	default:
		return NoTax
	}
}
