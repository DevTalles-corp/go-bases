package checkout

func TryChangeCustomerByValue(o Order, name string) {
	o.Customer = name
}

func ChangeCustomerByPointer(o *Order, name string) {
	o.Customer = name
}

func setCity(o *Order, city string) {
	if o.Meta == nil {
		o.Meta = map[string]string{}
	}
	o.Meta["city"] = city // Map, Slice, func, pointer, chan
}
