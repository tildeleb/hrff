// Copyright © 2012-2013 Lawrence E. Bakst. All rights reserved.

package siflags_test

import "siflags"
import "testing"

const default_i1 = 2
const default_i2 = 3

var _i1 siflags.Int64Value = default_i1
var i1 int64

const ci1 = 1 * 1024 * 1024 * 1024
const ci2 = -2

var _i2 siflags.Int64Value = default_i2
var i2 int64

const default_f1 = 4.4
const default_f2 = 5.5

var _f1 siflags.Float64Value = default_f1
var f1 float64
var _f2 siflags.Float64Value = default_f2
var f2 float64

func TestSIFlags(t *testing.T) {
	_i1.Set("1Gi")
	i1 = int64(_i1)
	if i1 != ci1 {
		t.Errorf("wanted %d got %d\n", ci1, i1)
	}
	_f1.Set("123.k")
	f1 = 123. * 1000
	if float64(_f1) != f1 {
		t.Errorf("wanted %f got %f\n", f1, _f1)
	}
	_f1.Set("-1.M")
	f1 = -1 * 1000000
	if float64(_f1) != f1 {
		t.Errorf("wanted %f got %f\n", f1, _f1)
	}
	_f1.Set(".001m")
	f1 = .000001
	if float64(_f1) != f1 {
		t.Errorf("wanted %f got %f\n", f1, _f1)
	}

}
