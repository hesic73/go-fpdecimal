package gofpdecimal

type errorString struct{ v string }

func (e *errorString) Error() string { return e.v }

var (
	errEmptyString            = &errorString{"empty string"}
	errMissingDigitsAfterSign = &errorString{"missing digits after sign"}
	errBadDigit               = &errorString{"bad digit"}
	errMultipleDots           = &errorString{"multiple dots"}
	errOverflow               = &errorString{"numeric overflow"}
	errDivisionByZero         = &errorString{"division by zero"}
)
