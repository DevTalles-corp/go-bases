package checkout

type Coupon struct {
	Code string
	Kind string
	Val  int
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
