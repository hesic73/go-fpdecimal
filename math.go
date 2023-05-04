package gofpdecimal

import (
	"math"

	"github.com/kpango/glg"
)

// Preserve the larger precision. It will throw an error when these efforts fail
func Add(a, b FpDecimal) (FpDecimal, error) {
	if a.precision < b.precision {
		a, b = b, a
	}
	sign0 := b.underlyingValue > 0
	new_b := b.underlyingValue * pow10(a.precision-b.precision)
	sign1 := new_b > 0

	if sign0 != sign1 {
		return GetZero(), errOverflow
	}
	tmp := a.underlyingValue + new_b
	if (tmp > a.underlyingValue) != (new_b > 0) {
		return GetZero(), errOverflow
	}
	p := a.precision
	for p > 0 && tmp%10 == 0 {
		p--
		tmp /= 10
	}
	return FpDecimal{
		underlyingValue: tmp,
		precision:       p,
	}, nil
}

// Preserve the larger precision. It will throw an error when these efforts fail
func Sub(a, b FpDecimal) (FpDecimal, error) {
	var tmp int64
	var p uint
	if a.precision < b.precision {
		sign0 := a.underlyingValue > 0
		new_a := a.underlyingValue * pow10(b.precision-a.precision)
		sign1 := new_a > 0
		if sign0 != sign1 {
			return GetZero(), errOverflow
		}

		tmp = new_a - b.underlyingValue
		if (tmp < new_a) != (b.underlyingValue > 0) {
			return GetZero(), errOverflow
		}
		p = b.precision

	} else {
		sign0 := b.underlyingValue > 0
		new_b := b.underlyingValue * pow10(a.precision-b.precision)
		sign1 := new_b > 0

		if sign0 != sign1 {
			return GetZero(), errOverflow
		}
		tmp = a.underlyingValue - new_b
		if (tmp < a.underlyingValue) != (new_b > 0) {
			return GetZero(), errOverflow
		}
		p = a.precision
	}

	for p > 0 && tmp%10 == 0 {
		p--
		tmp /= 10
	}
	return FpDecimal{
		underlyingValue: tmp,
		precision:       p,
	}, nil
}

func (d FpDecimal) MulInteger(i int64) (FpDecimal, error) {
	for i%10 == 0 && d.precision > 0 {
		i /= 10
		d.precision--
	}
	if d.underlyingValue*i/i != d.underlyingValue {
		return GetZero(), errOverflow
	}
	d.underlyingValue *= i
	return d, nil
}

// no guarantee on precision of the result
func (d FpDecimal) DivInteger(i int64) (FpDecimal, error) {
	if i == 0 {
		return GetZero(), errDivisionByZero
	}
	if d.underlyingValue == math.MinInt64 && i == -1 {
		return GetZero(), errOverflow
	}

	for i%10 == 0 && d.precision < 19 {
		i /= 10
		d.precision++
	}

	for d.precision < 19 && d.underlyingValue*10/10 == d.underlyingValue {
		d.underlyingValue *= 10
		d.precision++
	}

	d.underlyingValue /= i

	d.tight()
	return d, nil
}

// 最开始设计的时候就有问题，我把精度和magnitude混为一谈
func (d FpDecimal) Div(v FpDecimal) (FpDecimal, error) {

	if v.IsZero() {
		return GetZero(), errDivisionByZero
	}
	if d.underlyingValue == math.MinInt64 && v.underlyingValue == -1 {
		return GetZero(), errOverflow
	}

	// d.precision might be >=19 ?
	for v.underlyingValue%10 == 0 {
		v.underlyingValue /= 10
		d.precision++
	}
	for d.underlyingValue*10/10 == d.underlyingValue {
		d.underlyingValue *= 10
		d.precision++
	}

	if d.underlyingValue/v.underlyingValue < 1000000 {
		glg.Warn("FpDecimal.Div精度不够，转换为float64计算")
		tmp := FromFloat64(d.Float64()/v.Float64(), 18)
		return tmp, nil
	}

	d.underlyingValue /= v.underlyingValue
	if d.precision < v.precision {
		// 这里就属于设计缺陷，小数位和精度应该是两个不同的东西
		// 但改回int的话，之前好像默认precision非负，就先这样吧
		glg.Fatal("FpDecimal.Div d.precision<v.precision")
	}
	d.precision -= v.precision
	d.tight()
	return d, nil
}
