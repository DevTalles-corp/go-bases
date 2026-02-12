package checkout

type Money = int //Centavos

type Item struct {
	SKU   string
	Name  string
	Price Money
	Qty   int
}

type Order struct {
	ID       string
	Customer string
	Items    []Item
	Meta     map[string]string // City, state, source
}

type Totals struct {
	Subtotal Money
	Discount Money
	Tax      Money
	Shipping Money
	Total    Money
}
