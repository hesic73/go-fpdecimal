package gofpdecimal

import (
	"fmt"
	"log"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToLite(t *testing.T) {
	f := 40008.34
	tmp := fmt.Sprintf("%.04f", f)
	fmt.Println(tmp)
	a, err := FromString(tmp)
	if err != nil {
		log.Fatal(err)
	}
	a.To(2)
	b := FromFloat64(f, 2)
	if !assert.Equal(t, a.String(), b.String()) {
		fmt.Println(f, a, b)
	}
}

// As many methods are directly copied from fpdecimal, their tests are skipped
func TestTo(t *testing.T) {
	n := 10000
	for i := 0; i < n; i++ {
		f := float64(rand.Int63()%1000000000) / 10000.0
		tmp := fmt.Sprintf("%.04f", f)
		a, err := FromString(tmp)
		if err != nil {
			log.Fatal(err)
		}
		a.To(2)

		b := FromFloat64(f, 2)
		if !assert.Equal(t, a.String(), b.String()) {
			fmt.Println(i, f, a, b)
			break
		}
	}

}

func TestPercent(t *testing.T) {
	a, err := FromString("0.0000123009")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(a.Percent())
}
