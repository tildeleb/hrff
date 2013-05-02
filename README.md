#hrff

***WARNING: This project under heavy development right now with new capabilities on the way. The package name will be changing as will some of the data structures. Use at your own risk as the next version will be incompatible in some significant ways. -- leb***

Allows two capabilities.

1. SI unit prefixes to be used as multiplicative suffixes to numeric flags in Go Lang. 
For example you can use flags like the block size argument below:

		% # set block size to 1024
		% copy file1 file2 -blocksize=1Ki

2. Variables with one of the two new types can printed in human readable form, using the %[hH] format character, optionally with a unit string. Lower case 'h' is used to format decimal SI suffixes and upper case 'H" is used to format power of two SI suffixes. For example:

		var i1 Int64 = Int64{2000000, "bps"}
		var i2 Int64 = Int64{V: 3000}
		var f1 Float64 = Float64{-7020000000.1, "B"}
		var f2 Float64 = Float64{.000001, "s"}
		var imm1 Int64 = Int64{-40000, ""}
		var imm2 Int64 = Int64{1 * 1024 * 1024 * 1024, ""}
		fmt.Printf("i1=%d, i2=%d, f1=%f, f2=%f, i1=%h, i2=%h, f1=%0.4h, f2=%h, imm1=%h, imm2=%H\n",
			i1.V, i2.V, f1.V, f2.V, i1, i2, f1, f2, imm1, imm2)
		
		yields:
		i1=2000000, i2=3000, f1=-7020000000.100000, f2=0.000001, i1=2Mbps, i2=3k, f1=-7.0200GB, f2=1µs, imm1=-40k, imm2=1Gi

All standard SI unit prefixes are supported. In keeping with the SI standard "Ki" is used to multiply by 1024 and not "k" which is used to multiply by 1000.

This package extends the package "flags" by adding two new types and cooresponding Set(string) methods.


##The Two New Types
	Int64
	Float64

##Usage
	var blocksize hrff.Int64 = Int64{V: 512, U: ""}

	func main() {
		flag.Var(&blocksize, "blocksize", "blocksize for copy")
		flag.Parse()
		copy(src, dot, block size.V)
	}

##Bugs
The are a few conflicts with hex numbers. Is "0x1d" 30 or 1 deci? Similar conflicts occur with exa, deka, femto, and atto. I suspect hex numbers aren't very useful with multiplicative suffixes so the conflict is now resolved if favor of the suffix. I think that's acceptable. It could be fixed by allowing or even forcing a space between the number and the suffix. I'll take feedback on that issue. I'm considering optionally allowing the space.


