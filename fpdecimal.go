package gofpdecimal

import (
	"fmt"
	"math"
	"strings"
)

type FpDecimal struct {
	underlyingValue int64
	precision       uint // 名字好像起的比较失败，但又不是exp，应该说我被fpdecimal给带歪了。有空大改
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

func FromInt64(i int64) FpDecimal {
	return FpDecimal{
		underlyingValue: i,
		precision:       0,
	}
}

func (d *FpDecimal) tight() {
	for d.underlyingValue%10 == 0 && d.precision > 0 {
		d.underlyingValue /= 10
		if d.underlyingValue == 0 {
			d.precision = 0
			break
		}
		d.precision--
	}
}

func FromFloat64(f float64, precision uint) FpDecimal {
	v := floatPow10(f, precision)
	if v > math.MaxInt64 {
		panic("f*10^precision should be no larger than math.MaxInt64")
	} else if v < math.MinInt64 {
		panic("f*10^precision should be no less than math.MinInt64")
	}

	tmp := FpDecimal{
		underlyingValue: int64(v),
		precision:       precision,
	}

	tmp.tight()
	return tmp
}

func (d *FpDecimal) Float32() float32 {
	return float32(d.underlyingValue) / float32(pow10(d.precision))
}

func (d *FpDecimal) Float64() float64 {
	return float64(d.underlyingValue) / float64(pow10(d.precision))
}

func (d FpDecimal) String() string {
	return FixedPointDecimalToString(d.underlyingValue, int(d.precision))
}

func (d FpDecimal) Percent() string {
	if d.underlyingValue*100/100 != d.underlyingValue {
		fmt_str := fmt.Sprintf("%%.%df", d.precision)
		return fmt.Sprintf(fmt_str, d.Float64()*100) + "%"
	}
	d.underlyingValue *= 100
	d.tight()
	return FixedPointDecimalToString(d.underlyingValue, int(d.precision)) + "%"
}

func FromString(s string) (FpDecimal, error) {
	v, p, err := ParseFixedPointDecimal(s)
	return FpDecimal{
		underlyingValue: v,
		precision:       p,
	}, err
}

func (d *FpDecimal) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && len(b) >= 2 && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	d.underlyingValue, d.precision, err = ParseFixedPointDecimal(string(b))
	return err
}

func (d *FpDecimal) MarshalJSON() ([]byte, error) { return []byte(d.String()), nil }

func (d *FpDecimal) To(precision uint) bool {
	if precision > 18 {
		return false
	}
	if precision >= d.precision {
		return true
	}
	for i := precision; i < d.precision; i++ {
		d.underlyingValue /= 10
	}
	d.precision = precision
	d.tight()
	return true
}

func (d *FpDecimal) IsZero() bool {
	return d.underlyingValue == 0
}

func (d *FpDecimal) ToPrecision(n uint) string {
	if n >= d.precision {
		s := FixedPointDecimalToString(d.underlyingValue, int(d.precision))
		s = s + strings.Repeat("0", int(n-d.precision))
		return s
	} else {
		mask := pow10(d.precision - n)
		d.underlyingValue -= d.underlyingValue % mask
		d.tight()
		return FixedPointDecimalToString(d.underlyingValue, int(d.precision))
	}
}
