package math

func Calculator(a int, b int, op string) {
	if op == "+" {
		add(a, b)
	}
	if op == "-" {
		sub(a, b)
	}
}