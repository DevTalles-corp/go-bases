package checkout

import (
	"fmt"
	"strings"
)

func sign(neg bool) string {
	if neg {
		return "-"
	}

	return ""
}

func formatThousands(number int64) string {
	s := fmt.Sprintf("%d", number)
	if len(s) <= 3 {
		return s
	}

	var builder strings.Builder
	rem := len(s) % 3
	if rem == 0 {
		rem = 3
	}
	builder.WriteString(s[:rem])

	for i := rem; i < len(s); i += 3 {
		builder.WriteByte(',')
		builder.WriteString(s[i : i+3])
	}

	return builder.String()
}

func StringUSD(money Money) string {
	neg := money < 0
	if neg {
		money = -money
	}

	dollars := int64(money / 100)
	cents := int64(money % 100)

	return fmt.Sprintf("%s$%s.%02d USD", sign(neg), formatThousands(dollars), cents)
}
