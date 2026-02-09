package checkout

func RunDemo() {
	PrintHeader("Hola Checkout Engine :)")

	order := NewOrder("ORDER-001", "RICARDO")
	AddItem(&order, Item{SKU: "KB-001", Name: "Teclado", Price: 3500, Qty: 1})
	AddItem(&order, Item{SKU: "MB-024", Name: "Monitor", Price: 15000, Qty: 2})

	PrintKV("OrderID", order.ID)
	PrintKV("Customer", order.Customer)
	PrintKV("Items", len(order.Items))

	PrintDivider()
}
