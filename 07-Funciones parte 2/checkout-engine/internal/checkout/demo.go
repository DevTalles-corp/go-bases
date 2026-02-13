package checkout

import "fmt"

func RunDemo() {
	PrintHeader("Hola Checkout Engine :)")

	order := NewOrder("ORDER-001", "RICARDO")
	AddItem(&order, Item{SKU: "KB-001", Name: "Teclado", Price: 3500, Qty: 1})
	AddItem(&order, Item{SKU: "MB-024", Name: "Monitor", Price: 15000, Qty: 2})
	AddItem(&order, Item{SKU: "MB-054", Name: "CPU", Price: 45000, Qty: 3})

	//Probando validador
	PrintKV("VALIDADOR: ", ValidateOrder(order))

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

	PrintDivider()

	TryChangeCustomerByValue(order, "Nuevo nombre")
	PrintKV("Customer no cambia: ", order.Customer)

	ChangeCustomerByPointer(&order, "Andrei Cuéllar")
	PrintKV("Customer si cambia: ", order.Customer)

	setCity(&order, "Buenos Aires")
	PrintKV("Ciudad (map si cambia)", order.Meta["city"])

	PrintDivider()
	items := []Item{
		{SKU: "MS-003", Name: "Mouse", Price: 1200, Qty: 1},
		{SKU: "HD-005", Name: "HDMI", Price: 300, Qty: 2},
	}

	AddItems(&order, items...)
	PrintKV("Cantidad total: ", CalcTotalQty(order))
	PrintKV("Items: ", order.Items)

	PrintDivider()

	findItem, extraValueFind := FindItem(order, "MS-003")
	PrintKV2("Item encontrado", findItem, extraValueFind)
	getMeta, extraGetMeta := GetMeta(order, "city")
	PrintKV2("Metadato encontrado", getMeta, extraGetMeta)
	IndexOfItemValue, IndexOfItemExtra := IndexOfItem(order, "HD-008")
	PrintKV2("Index encontrado", IndexOfItemValue, IndexOfItemExtra)

	PrintDivider()

	couponValue, couponError := ParseCoupon("SAVE30")
	PrintKV2("Probando cupón: ", couponValue, couponError)

	PrintDivider()
	// computeValue, computeError := Compute(order)
	// _, computeError := Compute(order)
	// PrintKV2("Computar valores por nombre (TOTALES): ", computeValue, computeError)
	// PrintKV("TOTALES: ", computeError)

	PrintDivider()

	PrintKV("Descuento: ", FlatDiscount(200)(order))
	th := ThresholdPercentDiscount(2000, 20)
	PrintKV("Descuento %: ", th(order))

	PrintDivider()

	cityDiscount := func(order Order) Money {
		city, _ := GetMeta(order, "city")
		if city == "Buenos Aires" {
			return 200
		}
		return 0
	}

	PrintKV("Descuento especial por ciudad: ", cityDiscount(order))

	PrintDivider()

	discountKeyboard := MakeSKUDiscount("TECHProduct", 500)
	discountMouse := MakeSKUDiscount("HD-005", 150)

	fmt.Println(discountKeyboard(order))
	fmt.Println(discountMouse(order))

	PrintDivider()

	computeValue, computeError := Compute(order, NoTax, FreeShipping, FlatDiscount(5000), ThresholdPercentDiscount(2000, 10))
	PrintKV2("Computar valores por nombre (TOTALES): ", computeValue, computeError)
}
