package checkout

type DiscountFn func(Order) Money

type Coupon struct {
	Code string
	Kind string
	Val  int
}

type CompositeDiscount struct {
	Name string
	Fns  []DiscountFn
}

func ApplyCouponCodes(order *Order, codes ...string) {
	if order.Meta == nil {
		order.Meta = map[string]string{}
	}

	order.Meta["coupons"] = joinCoupons(codes) // slices de strings lo quiero pasar a string
	// ["EnvioGratis", "SANVALENTIN"] -> "EnvioGratis,SANVALENTIN"
}

func joinCoupons(coupons []string) string {
	if len(coupons) == 0 {
		return ""
	}

	out := coupons[0]
	for i := 1; i < len(coupons); i++ {
		out += "," + coupons[i]
	}

	return out
}

func FlatDiscount(amount Money) DiscountFn {
	return func(order Order) Money {
		return amount
	}
}

func ThresholdPercentDiscount(min Money, percent int) DiscountFn {
	return func(order Order) Money {
		sub := order.CalcSubtotal()
		if sub < min {
			return 0
		}
		return sub * Money(percent) / 100
	}
}

func MakeSKUDiscount(sku string, amount Money) DiscountFn {
	return func(order Order) Money {
		_, ok := order.FindItem(sku)
		if !ok {
			return 0
		}

		return amount
	}
}

func ApplyDiscountsRecursive(order Order, fns []DiscountFn) Money {
	if len(fns) == 0 {
		return 0
	}

	return fns[0](order) + ApplyDiscountsRecursive(order, fns[1:]) // "FREESHIPPING", "CITYDISCOUNT", "HOTSALE123"
}

func (composite CompositeDiscount) Apply(order Order) Money {
	return ApplyDiscountsRecursive(order, composite.Fns)
}
