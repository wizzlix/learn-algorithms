package algorithms

func Devisor(a, b int) int {
	if a == 0 {
		return b
	} else if b == 0 {
		return a
	} else if a > b {
		return Devisor(a%b, b)
	} else {
		return Devisor(a, b%a)
	}
}
