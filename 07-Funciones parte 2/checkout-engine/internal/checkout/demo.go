package checkout

import "fmt"

func RunDemo() {
	PrintHeader("Hola Checkout Engine :)")

	order := NewOrder("ORDER-001", "RICARDO")
	order.AddItem(Item{SKU: "MB-024", Name: "Monitor", Price: 15000, Qty: 2})
	order.AddItem(Item{SKU: "KB-001", Name: "Teclado", Price: 3500, Qty: 1})
	order.AddItem(Item{SKU: "MB-054", Name: "CPU", Price: 45000, Qty: 3})

	//Probando validador
	PrintKV("VALIDADOR: ", ValidateOrder(order))

	PrintKV("OrderID", order.ID)
	PrintKV("Customer", order.Customer)
	PrintKV("Items", len(order.Items))

	remove := order.RemoveItem("KB-001")
	PrintKV("Removed KB-001: ", remove)

	PrintDivider()

	sub := order.CalcSubtotal()
	qty := order.CalcTotalQty()

	PrintKV("Subtotal: ", StringUSD(sub))
	PrintKV("Cantidad: ", qty)

	PrintDivider()

	TryChangeCustomerByValue(order, "Nuevo nombre")
	PrintKV("Customer no cambia: ", order.Customer)

	ChangeCustomerByPointer(&order, "Andrei Cuéllar")
	PrintKV("Customer si cambia: ", order.Customer)

	setCity(&order, "NL")
	PrintKV("Ciudad (map si cambia)", order.Meta["city"])

	//Setear la zona
	setZone(&order, "NATIONAL")

	PrintDivider()
	items := []Item{
		{SKU: "MS-003", Name: "Mouse", Price: 1200, Qty: 1},
		{SKU: "HD-005", Name: "HDMI", Price: 300, Qty: 2},
	}

	order.AddItems(items...)
	PrintKV("Cantidad total: ", order.CalcTotalQty())
	PrintKV("Items: ", order.Items)

	PrintDivider()

	findItem, extraValueFind := order.FindItem("MS-003")
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

	PrintKV("Descuento: ", StringUSD(FlatDiscount(200)(order)))
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

	state, _ := GetMeta(order, "city")
	zone, _ := GetMeta(order, "zone")

	taxFn := NewTaxByState(state)
	shipFn := NewShippingByZone(zone)

	promo := CompositeDiscount{
		Name: "Promoción Febrero",
		Fns:  []DiscountFn{
			// FlatDiscount(100),
			// ThresholdPercentDiscount(2000, 10),
			// MakeSKUDiscount("HD-005", 200),
		},
	}

	bundle := promo.Apply(order)
	PrintKV("DESCUENTO RECURSIVO: ", StringUSD(bundle))

	computeValue, computeError := Compute(order, bundle, taxFn, shipFn, FlatDiscount(5000), ThresholdPercentDiscount(2000, 10))
	PrintKV2("Computar valores por nombre (TOTALES): ", computeValue, computeError)
	PrintKV("TOTAL: ", StringUSD(computeValue.Total))
}
