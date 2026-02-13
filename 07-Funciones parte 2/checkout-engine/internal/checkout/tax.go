package checkout

type TaxFn func(Order) Money

func NoTax(Order) Money {
	return 0
}

func IVA16(order Order) Money {
	sub := CalcSubtotal(order)
	return sub * 16 / 100
}
