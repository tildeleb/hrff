// Copyright © 2012-2013 Lawrence E. Bakst. All rights reserved.

package main

import "github.com/tildeleb/hrff"
import "fmt"
import "flag"

var i1, i2 hrff.Int64
var f1, f2 hrff.Float64

func main() {
	i1.Set("123Ki")
	i2 = hrff.Int64{456 * 1024 * 1024 * 1024, "B"}
	f1.Set("2.345G")
	f2 = hrff.Float64{V: 1.234 * 1024 * 1024 * 1024 * 1024}
	flag.Var(&i1, "i1", "i1")
	flag.Var(&i2, "i2", "i2")
	flag.Var(&f1, "f1", "f1")
	flag.Var(&f2, "f2", "f2")

	flag.Parse()
	fmt.Printf("i1=%H, i2=%H, f1=%0.3h, f2=%0.0h\n", i1, i2, f1, f2)
}
