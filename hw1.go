package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	var in1, in2, in3 string
	var a, b, c float64

	for {
		n, _ := fmt.Scan(&in1, &in2, &in3)
		if n <= 0 {
			break
		}
		if n < 3 {
			fmt.Println("please input three number")
			continue
		}
		var err1, err2, err3 error
		a, err1 = strconv.ParseFloat(in1, 64)
		b, err2 = strconv.ParseFloat(in2, 64)
		c, err3 = strconv.ParseFloat(in3, 64)
		if err1 != nil || err2 != nil || err3 != nil {
			fmt.Println("please input three number")
			continue
		}
		var area, s float64
		s = (a + b + c) / 2
		area = math.Sqrt(s * (s - a) * (s - b) * (s - c))
		fmt.Println("the area of traingle: ", a, b, c, "is", area)
	}
}

//https://yourbasic.org/golang/convert-string-to-float/
//https://pkg.go.dev/search?q=fmt.Scan&m=
