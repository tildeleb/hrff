// Copyright © 2012-2013 Lawrence E. Bakst. All rights reserved.

package siflags

import "fmt"
import "strconv"

// Pakcage allows command line arguments like % dd bs=1Mi
// defines two news types, Int64Value and Float64Value which provide methods for flags to accept these kind of args
// My usual approach is to convert them to a build-in int/float type after flags has parsed the command line
// If you want to use "K", "M", "G", "T", "P", "E" instead of Ki", "Mi", "Gi", "Ti", "Pi", "Ei" please call Wrong()

// yes I know about iota but it doesn't really work here and I find what's below clearer even if it did
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

type Int64Value int64
type Float64Value float64

// thanks to my early mentor
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

func (r *Int64Value) Set(s string) error {

	m, l, _ := getPrefix(s)
	v, err := strconv.ParseInt(s[:l], 10, 64)
	if err != nil {
		return err
	}
	// fmt.Printf("Set: v=%d, m=%f, v*m=%v\n", v, m, v*int64(m))
	*r = Int64Value(v * int64(m))
	return err
}

func (r *Float64Value) Set(s string) error {

	m, l, ok := getPrefix(s)
	v, er := strconv.ParseFloat(s[:l], 64)
	if !ok {
		return er
	}
	// fmt.Printf("Set: v=%f, m=%f, v*m=%v\n", v, m, v*m)
	*r = Float64Value(v * m)
	return er
}

// maybe someday these will actually print numbers out using SI unit prefixes
// maybe delete for now?
func (v *Int64Value) String() string {
	return fmt.Sprintf("%d", *v)
}

func (v *Float64Value) String() string {
	return fmt.Sprintf("%f", *v)
}
