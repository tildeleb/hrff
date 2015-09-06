// Copyright © 2012-2014 Lawrence E. Bakst. All rights reserved.

package main

import "leb.io/hrff"
import "fmt"
import "flag"

var i1, i2 hrff.Int64
var f1, f2 hrff.Float64

func main() {
	//fmt.Printf("i1=%h i1=%H\n", i1, i1)
	i1.Set("123Ki")
	i2 = hrff.Int64{456 * 1024 * 1024 * 1024, "B"}
	f1.Set("2.345G")
	f2 = hrff.Float64{V: 1.234 * 1024 * 1024 * 1024 * 1024}
	flag.Var(&i1, "i1", "i1")
	flag.Var(&i2, "i2", "i2")
	flag.Var(&f1, "f1", "f1")
	flag.Var(&f2, "f2", "f2")

	flag.Parse()
	//fmt.Printf("i1=%H, i2=%H, f1=%0.3h, f2=%0.0h\n", i1, i2, f1, f2)

	iv := hrff.Int64{9, "B"}

	for i := 0; i < 19; i++ {
		fmt.Printf("%d, %h\n", iv.V, iv)
		iv.V *= 10
	}
	fmt.Printf("---\n")
	fv := hrff.Float64{0.9, "l"}
	for i := 0; i < 25; i++ {
		fmt.Printf("%.25f, %h\n", fv.V, fv)
		fv.V /= 10.0
	}
	fmt.Printf("---\n")
	fv.V = 9.0;
	for i := 0; i < 34; i++ {
		fmt.Printf("%.31f, %h\n", fv.V, fv)
		fv.V *= 10.0
	}
}
