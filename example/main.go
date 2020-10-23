package main

import (
	"fmt"
	"time"

	"github.com/ace-teknologi/isdst"
)

func main() {
	loc, _ := time.LoadLocation("Australia/Broken_Hill")

	// DST
	t := time.Date(2020, time.January, 1, 0, 0, 0, 0, loc)
	fmt.Printf("Is daylight savings? %t", isdst.IsDST(t))

	// Non-DST
	t = time.Date(2020, time.June, 1, 0, 0, 0, 0, loc)
	fmt.Printf("Is daylight savings? %t", isdst.IsDST(t))
}
