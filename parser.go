package gofpdecimal

const sep = '.'

func ParseFixedPointDecimal(s string) (int64, uint, error) {
	if s == "" {
		return 0, 0, errEmptyString
	}

	s0 := s
	if s[0] == '-' || s[0] == '+' {
		s = s[1:]
		if len(s) < 1 {
			return 0, 0, errMissingDigitsAfterSign
		}
	}

	var d int8 = -1 // current decimal position
	var n int64     // output
	for _, ch := range []byte(s) {
		if d == 18 {
			break
		}

		if ch == sep {
			if d != -1 {
				return 0, 0, errMultipleDots
			}
			d = 0
			continue
		}

		ch -= '0'
		if ch > 9 {
			return 0, 0, errBadDigit
		}
		n = n*10 + int64(ch)
		if n < 0 {
			return 0, 0, errOverflow
		}

		if d != -1 {
			d++
		}
	}

	// fill rest of 0
	if d == -1 {
		d = 0
	}

	if s0[0] == '-' {
		n = -n
	}

	for n%10 == 0 && d > 0 {
		n /= 10
		d--
	}

	return n, uint(d), nil
}
