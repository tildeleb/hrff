// Copyright © 2012-2013 Lawrence E. Bakst. All rights reserved.

package main

import "siflags"
import "fmt"
import "flag"

var tf = flag.Bool("t", false, "test flag")

const default_i1 = 2
const default_i2 = 3

var _i1 siflags.Int64Value = default_i1
var i1 int64
var _i2 siflags.Int64Value = default_i2
var i2 int64

const default_f1 = 4.4
const default_f2 = 5.5

var _f1 siflags.Float64Value = default_f1
var f1 float64
var _f2 siflags.Float64Value = default_f2
var f2 float64

// 6.out -i1=1Ki -i2=2Gi -f1=1H -f2=6m
func main() {
	flag.Var(&_i1, "i1", "i1")
	flag.Var(&_i2, "i2", "i2")
	flag.Var(&_f1, "f1", "f1")
	flag.Var(&_f2, "f2", "f2")

	flag.Parse()
	switch {
	case *tf:
		fmt.Printf("test\n")
	default:
		i1 = int64(_i1)
		i2 = int64(_i2)
		f1 = float64(_f1)
		f2 = float64(_f2)
		fmt.Printf("i1=%d, i2=%d, f1=%f, f2=%f\n", i1, i2, f1, f2)
	}
}
