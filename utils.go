package gofpdecimal

func pow10(e uint) int64 {
	if e == 0 {
		return 1
	} else if e == 1 {
		return 10
	} else if e == 2 {
		return 100
	} else if e == 3 {
		return 1000
	} else if e == 4 {
		return 10000
	} else if e == 5 {
		return 100000
	} else if e == 6 {
		return 1000000
	} else if e == 7 {
		return 10000000
	} else if e == 8 {
		return 100000000
	} else if e == 9 {
		return 1000000000
	} else if e == 10 {
		return 10000000000
	} else if e == 11 {
		return 100000000000
	} else if e == 12 {
		return 1000000000000
	} else if e == 13 {
		return 10000000000000
	} else if e == 14 {
		return 100000000000000
	} else if e == 15 {
		return 1000000000000000
	} else if e == 16 {
		return 10000000000000000
	} else if e == 17 {
		return 100000000000000000
	} else if e == 18 {
		return 1000000000000000000
	}
	panic("The precision should be at most 18")
}

func pow10f(e uint) float64 {
	if e == 0 {
		return 1
	} else if e == 1 {
		return 10
	} else if e == 2 {
		return 100
	} else if e == 3 {
		return 1000
	} else if e == 4 {
		return 10000
	} else if e == 5 {
		return 100000
	} else if e == 6 {
		return 1000000
	} else if e == 7 {
		return 10000000
	} else if e == 8 {
		return 100000000
	} else if e == 9 {
		return 1000000000
	} else if e == 10 {
		return 10000000000
	} else if e == 11 {
		return 100000000000
	} else if e == 12 {
		return 1000000000000
	} else if e == 13 {
		return 10000000000000
	} else if e == 14 {
		return 100000000000000
	} else if e == 15 {
		return 1000000000000000
	} else if e == 16 {
		return 10000000000000000
	} else if e == 17 {
		return 100000000000000000
	} else if e == 18 {
		return 1000000000000000000
	}
	panic("The precision should be at most 18")
}

func floatPow10(f float64, e uint) float64 {
	if e > 18 {
		panic("The precision should be at most 18")
	}
	f *= pow10f(e)
	return f + 1e-8 // or .99999999 will be rounded
}

func log10Int(v int64) int {
	// 在我关心的case中当成1好像无所谓
	if v == 0 {
		panic("log_10^0")
	}
	if v < 10 {
		return 1
	} else if v < 100 {
		return 2
	} else if v < 1000 {
		return 3
	} else if v < 10000 {
		return 4
	} else if v < 100000 {
		return 5
	} else if v < 1000000 {
		return 6
	} else if v < 10000000 {
		return 7
	} else if v < 100000000 {
		return 8
	} else if v < 1000000000 {
		return 9
	} else if v < 10000000000 {
		return 10
	} else if v < 100000000000 {
		return 11
	} else if v < 1000000000000 {
		return 12
	} else if v < 10000000000000 {
		return 13
	} else if v < 100000000000000 {
		return 14
	} else if v < 1000000000000000 {
		return 15
	} else if v < 10000000000000000 {
		return 16
	} else if v < 100000000000000000 {
		return 17
	} else if v < 1000000000000000000 {
		return 18
	} else {
		return 19
	}
}
