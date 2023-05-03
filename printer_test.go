package gofpdecimal_test

import (
	"fmt"
	"log"
	"testing"

	gofpdecimal "github.com/hesic73/go-fpdecimal"
)

func TestToPrecision(t *testing.T) {
	a, err := gofpdecimal.FromString("123.56789")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(a.ToPrecision(3))
}
