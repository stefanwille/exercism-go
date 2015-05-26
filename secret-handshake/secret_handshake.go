package secret

const (
	Wink int = 1 << iota
	DoubleBlink
	CloseEyes
	Jump
	Reverse
)

func Handshake(n int) []string {
	if n < 0 || n > 31 {
		n = 0
	}
	result := []string{}
	if n&Wink != 0 {
		result = append(result, "wink")
	}
	if n&DoubleBlink != 0 {
		result = append(result, "double blink")
	}
	if n&CloseEyes != 0 {
		result = append(result, "close your eyes")
	}
	if n&Jump != 0 {
		result = append(result, "jump")
	}
	if n&Reverse != 0 {
		result = reverse(result)
	}
	return result
}

func reverse(input []string) []string {
	if len(input) == 0 {
		return input
	}
	return append(reverse(input[1:]), input[0])
}
