package main

import (
	"fmt"
	"math"

	"github.com/niklastomas/go_crash_course/03_packages/strutil"
)

func main() {
	fmt.Println(math.Floor(3.2))
	fmt.Println(math.Ceil(3.8))
	fmt.Println(math.Sqrt(16))
	fmt.Println(strutil.Reverse("elloh"))

}
