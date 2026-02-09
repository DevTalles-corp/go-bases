package checkout

func RunDemo() {
	PrintHeader("Hola Checkout Engine :)")

	order := NewOrder("ORDER-001", "RICARDO")
	AddItem(&order, Item{SKU: "KB-001", Name: "Teclado", Price: 3500, Qty: 1})
	AddItem(&order, Item{SKU: "MB-024", Name: "Monitor", Price: 15000, Qty: 2})
	AddItem(&order, Item{SKU: "MB-054", Name: "CPU", Price: 45000, Qty: 3})

	PrintKV("OrderID", order.ID)
	PrintKV("Customer", order.Customer)
	PrintKV("Items", len(order.Items))

	remove := RemoveItem(&order, "KB-001")
	PrintKV("Removed KB-001: ", remove)

	PrintDivider()

	sub := CalcSubtotal(order)
	qty := CalcTotalQty(order)

	PrintKV("Subtotal: ", sub)
	PrintKV("Cantidad: ", qty)

}
