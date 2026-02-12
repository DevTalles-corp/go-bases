package checkout

type TaxFn func(Order) Money

func NoTax(Order) Money {
	return 0
}
