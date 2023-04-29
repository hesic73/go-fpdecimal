package gofpdecimal

import (
	"fmt"
	"log"
	"testing"
)

func TestFrom(t *testing.T) {
	d, err := FromString("323.15932906")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(d)
	fmt.Println(d.String())
	fmt.Println(d.Float64())
}
