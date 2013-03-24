package siflags_test

import "siflags"
import "testing"

const default_i1 = 2
const default_i2 = 3

var _i1 Int64Value = default_i1
var i1 int64

const ci1 = 1 * 1024 * 1024 * 1024

var _i2 Int64Value = default_i2
var i2 int64

const default_f1 = 4.4
const default_f2 = 5.5

var _f1 Float64Value = default_f1
var f1 float64
var _f2 Float64Value = default_f2
var f2 float64

func TestSIFlags(t *testing.T) {
	_i1.Set("1Gi")
	i1 = int64(_i1)
	if i1 != ci1 {
		t.Errorf("wanted %d got %d\n", ci1, i1)
	}
}
