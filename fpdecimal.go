package gofpdecimal

import (
	"math"
)

type FpDecimal struct {
	underlyingValue int64
	precision       uint
}

func GetZero() FpDecimal {
	return FpDecimal{
		underlyingValue: 0,
		precision:       0,
	}
}

func GetOne() FpDecimal {
	return FpDecimal{
		underlyingValue: 1,
		precision:       0,
	}
}

func (d *FpDecimal) GetPrecision() uint {
	return d.precision
}

func NewFpDecimal(precision uint) FpDecimal {
	return FpDecimal{
		underlyingValue: 0,
		precision:       precision,
	}
}

func FromInteger(i int64) FpDecimal {
	return FpDecimal{
		underlyingValue: int64(i),
		precision:       0,
	}
}

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

func FromFloat64(f float64, precision uint) FpDecimal {
	if f*float64(pow10(precision)) > math.MaxInt64 {
		panic("f*10^precision should be no larger than math.MaxInt64")
	} else if f*float64(pow10(precision)) < math.MinInt64 {
		panic("f*10^precision should be no less than math.MinInt64")
	}

	tmp := FpDecimal{
		underlyingValue: int64(f * float64(pow10(precision))),
		precision:       precision,
	}

	for tmp.underlyingValue%10 == 0 && tmp.precision > 0 {
		tmp.underlyingValue /= 10
		tmp.precision--
	}
	return tmp
}

func (d *FpDecimal) Float32() float32 {
	return float32(d.underlyingValue) / float32(pow10(d.precision))
}

func (d *FpDecimal) Float64() float64 {
	return float64(d.underlyingValue) / float64(pow10(d.precision))
}

func (d *FpDecimal) String() string {
	return FixedPointDecimalToString(d.underlyingValue, int(d.precision))
}

func FromString(s string) (FpDecimal, error) {
	v, p, err := ParseFixedPointDecimal(s)
	return FpDecimal{
		underlyingValue: v,
		precision:       p,
	}, err
}

func (d *FpDecimal) UnmarshalJSON(b []byte) (err error) {
	d.underlyingValue, d.precision, err = ParseFixedPointDecimal(string(b))
	return err
}

func (d *FpDecimal) MarshalJSON() ([]byte, error) { return []byte(d.String()), nil }
