package gofpdecimal_test

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"testing"

	gofpdecimal "github.com/hesic73/go-fpdecimal"
	"github.com/stretchr/testify/assert"
)

func generateDatapointFloat(n int, op int, ch chan []float64) {
	for i := 0; i < n; i++ {
		a := float64(rand.Int63()%1000000000) / 100000.0
		b := float64(rand.Int63()%1000000000) / 100000.0
		var c float64
		switch op {
		case 0:
			c = a + b
		case 1:
			c = a - b
		default:
			panic("unsupported op")

		}
		ch <- []float64{a, b, c}
	}
}

func TestAdd(t *testing.T) {
	ch := make(chan []float64)
	n := 10000
	go generateDatapointFloat(n, 0, ch)

	for i := 0; i < n; i++ {
		tmp := <-ch
		a := gofpdecimal.FromFloat64(tmp[0], 10)
		b := gofpdecimal.FromFloat64(tmp[1], 10)
		c := gofpdecimal.FromFloat64(tmp[2], 10)
		d, err := gofpdecimal.Add(a, b)
		if err != nil {
			log.Fatal(err)
		}

		if !assert.Equal(t, true, math.Abs(d.Float64()-tmp[2]) < 1e-8) {
			fmt.Println(a.String(), b.String(), c.String(), d.String())
			break
		}
	}
}

func TestSub(t *testing.T) {
	ch := make(chan []float64)
	n := 10000
	go generateDatapointFloat(n, 1, ch)

	for i := 0; i < n; i++ {
		tmp := <-ch
		a := gofpdecimal.FromFloat64(tmp[0], 10)
		b := gofpdecimal.FromFloat64(tmp[1], 10)
		c := gofpdecimal.FromFloat64(tmp[2], 10)
		d, err := gofpdecimal.Sub(a, b)
		if err != nil {
			log.Fatal(err)
		}

		if !assert.Equal(t, true, math.Abs(d.Float64()-tmp[2]) < 1e-8) {
			fmt.Println(a.String(), b.String(), c.String(), d.String())
			break
		}
	}
}

func TestMulInteger(t *testing.T) {
	a, err := gofpdecimal.FromString("12345678901234")
	if err != nil {
		log.Fatal(err)
	}
	b, err := a.MulInteger(88)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(a, b)
}

func TestDivInteger(t *testing.T) {
	a, err := gofpdecimal.FromString("12345678901234.123")
	if err != nil {
		log.Fatal(err)
	}
	b, err := a.DivInteger(88)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(a, b)
}
