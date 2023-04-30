package gofpdecimal

func pow10(e uint) int64 {
	if e > 18 {
		panic("The precision should be at most 18")
	}
	var tmp int64 = 1
	for e > 0 {
		tmp *= 10
		e--
	}
	return tmp
}

func floatPow10(f float64, e uint) float64 {
	if e > 18 {
		panic("The precision should be at most 18")
	}
	for e > 0 {
		f *= 10.0
		e--
	}
	return f + 1e-8 // or .99999999 will be rounded
}
