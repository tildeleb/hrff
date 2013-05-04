// Copyright © 2012-2013 Lawrence E. Bakst. All rights reserved.

// package main
package hrff

import "fmt"
import "strconv"

// import "flag"

// Package hrff (Human Readbale Flags and Formatting)
// Allows command line arguments like % dd bs=1Mi
// Defines two news types, Int64 and Float64 which provide methods for flags to accept these kind of args
// My usual approach is to convert them to a build-in int/float type after flags has parsed the command line
// If you want to use "K", "M", "G", "T", "P", "E" instead of Ki", "Mi", "Gi", "Ti", "Pi", "Ei" please call Classic()

// yes I know about iota but it doesn't really work here and I find what's below clearer
var SIsufixes map[string]float64 = map[string]float64{
	"H": 1000000000000000000000000000, // hella (one for the team)

	"Y":  1000000000000000000000000, // yota
	"Z":  1000000000000000000000,    // zetta
	"E":  1000000000000000000,       // exa
	"P":  1000000000000000,          // peta
	"T":  1000000000000,             // tera
	"G":  1000000000,                // giga
	"M":  1000000,                   // mega
	"k":  1000,                      // kilo
	"h":  100,                       // hecto
	"da": 10,                        // deka
	"":   1,                         // not real dummy stopper
	"d":  .1,                        // deci
	"c":  .01,                       // centi
	"m":  .001,                      // milli
	"µ":  .000001,                   // micro (unicode char see below)
	"n":  .000000001,                // nano
	"p":  .00000000001,              // pico
	"f":  .000000000000001,          // femto
	"a":  .000000000000000001,       // atto
	"z":  .000000000000000000001,    // zepto
	"y":  .000000000000000000000001, // yocto

	"u": .000001, // micro (with u)

	"Ki": 1024,                                                  // kibi	
	"Mi": 1024 * 1024,                                           // mebi	
	"Gi": 1024 * 1024 * 1024,                                    // gibi	
	"Ti": 1024 * 1024 * 1024 * 1024,                             // tebi	
	"Pi": 1024 * 1024 * 1024 * 1024 * 1024,                      // pebi	
	"Ei": 1024 * 1024 * 1024 * 1024 * 1024 * 1024,               // exbi
	"Zi": 1024 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024,        // zebi
	"Yi": 1024 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024, // yobi
}

var order []string = []string{"H", "Y", "Z", "E", "P", "T", "G", "M", "k", "h", "da", "", "d", "c", "m", "µ", "n", "p", "f", "a", "z", "y"}
var order2 []string = []string{"Yi", "Zi", "Ei", "Pi", "Ti", "Gi", "Mi", "Ki", "", "d", "c", "m", "µ", "n", "p", "f", "a", "z", "y"}

// considering removing this
func Classic() {
	SIsufixes["K"] = SIsufixes["Ki"]
	SIsufixes["M"] = SIsufixes["Mi"]
	SIsufixes["G"] = SIsufixes["Gi"]
	SIsufixes["T"] = SIsufixes["Ti"]
	SIsufixes["P"] = SIsufixes["Pi"]
	SIsufixes["E"] = SIsufixes["Ei"]
	SIsufixes["Z"] = SIsufixes["Zi"]
	SIsufixes["Y"] = SIsufixes["Yi"]
}

type Int64 struct {
	V int64
	U string
}
type Float64 struct {
	V float64
	U string
}

// thanks to my mentor
func knot(c rune, chars string) bool {
	for _, v := range chars {
		if c == v {
			return false
		}
	}
	return true
}

func getPrefix(s string) (float64, int, bool) {
	var m float64 = 1
	var o int = 0

	//	fmt.Printf("getPrefix: s=%q\n", s)
	_, ok := SIsufixes["xxx"] // better way?
	l := len(s)
	if l > 1 {
		if knot(rune(s[l-1]), "0123456789.") {
			if l > 2 {
				if knot(rune(s[l-2]), "0123456789.+-eE") {
					o = 2
				} else {
					o = 1
				}
			} else {
				o = 1
			}
		}
		m, ok = SIsufixes[s[l-o:]]
		//		fmt.Printf("getPrefix: %q, m=%f, l=%d, o=%d, ok=%v\n", s[l-o:], m, l, o, ok)
	}
	return m, l - o, ok
}

func pif(val int64, units string, p, w int, order []string) string {
	var sip string

	// fmt.Printf("pif: %d\n", val)
	sgn := ""
	if val < 0 {
		sgn = "-"
		val = -val
	}

	fs := fmt.Sprintf("%%s%%%d.%dd%%s%%s", w, p)
	// fmt.Printf("sgn=%q, fs=%q\n", sgn, fs)

	for _, sip = range order {
		//		fmt.Printf("Format: try %q, ", sip)
		if (SIsufixes[sip] <= float64(val)) || (sip == "" && val == 0) {
			break
		}
	}
	// fmt.Printf("pif: sip=%q\n", sip)
	val = val / int64(SIsufixes[sip])
	return fmt.Sprintf(fs, sgn, val, sip, units)
}

func pff(val float64, units string, p, w int, order []string) string {
	var sip string

	// fmt.Printf("pff: %f\n", val)
	sgn := ""
	if val < 0 {
		sgn = "-"
		val = -val
	}

	fs := fmt.Sprintf("%%s%%%d.%df%%s%%s", w, p)
	// fmt.Printf("sgn=%q, fs=%q\n", sgn, fs)

	for _, sip = range order {
		//		fmt.Printf("pff: %q, %f <= %f\n", sip, SIsufixes[sip], val)
		if SIsufixes[sip] == 1 {
			continue
		}
		if SIsufixes[sip] <= val {
			break
		}
	}
	// fmt.Printf("pff: i=%d, sip=%q\n", i, sip)
	val = val / SIsufixes[sip]
	str := fmt.Sprintf(fs, sgn, val, sip, units)
	return str
}

func i(v *Int64, s fmt.State, c rune) {
	var val int64 = int64(v.V)
	var str string

	w, wok := s.Width()
	p, pok := s.Precision()
	// fmt.Printf("i: c=%c, w=%d, wok=%v, p=%d, pok=%v\n", c, w, wok, p, pok)
	if !wok || !pok {
		str = ""
	}
	switch c {
	case 'h':
		str = pif(val, v.U, p, w, order)
	case 'H':
		str = pif(val, v.U, p, w, order2)
	case 'd':
		str = fmt.Sprintf("%d", val)
	case 'v':
		str = fmt.Sprintf("%v", val)
	default:
		fmt.Printf("default\n")
		str = fmt.Sprintf("%d", val)
	}
	b := []byte(str)
	s.Write(b)
}

func f(v *Float64, s fmt.State, c rune) {
	var val float64 = float64(v.V)
	var str string

	w, wok := s.Width()
	p, pok := s.Precision()
	// fmt.Printf("f: c=%c, w=%d, wok=%v, p=%d, pok=%v\n", c, w, wok, p, pok)
	if !wok || !pok {
		str = ""
	}
	switch c {
	case 'h':
		str = pff(val, v.U, p, w, order)
	case 'H':
		str = pff(val, v.U, p, w, order2)
	case 'd':
		str = fmt.Sprintf("%d", val)
	case 'v':
		str = fmt.Sprintf("%v", val)
	default:
		fmt.Printf("default\n")
		str = fmt.Sprintf("%d", val)
	}
	b := []byte(str)
	s.Write(b)
}

// FIX FIX FIX check ok or err? if no prefix we must convert anyway not err
func (r *Int64) Set(s string) error {

	m, l, _ := getPrefix(s)
	v, err := strconv.ParseInt(s[:l], 10, 64)
	if err != nil {
		return err
	}
	// fmt.Printf("Set: v=%d, m=%f, v*m=%v\n", v, m, v*int64(m))
	r.V = int64(v * int64(m))
	return err
}

func (r *Float64) Set(s string) error {

	m, l, ok := getPrefix(s)
	v, err := strconv.ParseFloat(s[:l], 64)
	if !ok {
		return err
	}
	// fmt.Printf("Set: v=%f, m=%f, v*m=%v\n", v, m, v*m)
	r.V = float64(v * m)
	return err
}

func (v Int64) String() string {
	//	fmt.Printf("String: I\n")
	return fmt.Sprintf("%s", pif(v.V, v.U, 0, 0, order))
}

func (v Float64) String() string {
	//	fmt.Printf("String: F\n")
	return fmt.Sprintf("%s", pff(v.V, v.U, 0, 0, order))
}

func (v Int64) Format(s fmt.State, c rune) {
	i(&v, s, c)
}

func (v Float64) Format(s fmt.State, c rune) {
	f(&v, s, c)
}

/*
var tf = flag.Bool("t", false, "test flag")

var i1 Int64 = Int64{2000000, "bps"}
var i2 Int64 = Int64{V: 3000}
var f1 Float64 = Float64{-7020000000.1, "B"}
var f2 Float64 = Float64{.000001, "s"}

var imm1 Int64 = Int64{-40000, ""}
var imm2 Int64 = Int64{1 * 1024 * 1024 * 1024, ""}

// 6.out -i1=1Ki -i2=2Gi -f1=1H -f2=6m
func main() {
	flag.Var(&i1, "10", "i1")
	flag.Var(&i2, "11", "i2")
	flag.Var(&f1, "12.1", "f1")
	flag.Var(&f2, "12.2", "f2")

	flag.Parse()
	switch {
	case *tf:
		fmt.Printf("test\n")
	default:
		fmt.Printf("i1=%d, i2=%d, f1=%f, f2=%f, i1=%h, i2=%h, f1=%0.4h, f2=%h, imm1=%h, imm2=%H\n",
			i1.V, i2.V, f1.V, f2.V, i1, i2, f1, f2, imm1, imm2)
		fmt.Printf("%10.3h\n", Int64{V: 0, U: "foobars"})
	}
}
*/
