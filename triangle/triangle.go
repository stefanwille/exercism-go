package triangle

type Kind string

const (
	Equ = "Equilateral"
	Iso = "Isosceles"
	Sca = "Scalene"
	NaT = "Error"
)

func KindFromSides(a, b, c float64) Kind {
	switch {
	case !(a > 0 && b > 0 && c > 0):
		return NaT
	case a+b <= c || a+c <= b || b+c <= a:
		return NaT
	case a == b && b == c:
		return Equ
	case a == b || b == c || c == a:
		return Iso
	default:
		return Sca
	}
}
