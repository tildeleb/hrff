#Human Readable Flags and Formatting

Provides two capabilities:

1. Human readable command line arguments.
SI unit prefixes can be used as multiplicative suffixes to numeric flags on the Go command line CLI. 
For example you can use flags like the block size argument below:

		% # set block size to 1024
		% copy file1 file2 -blocksize=1Ki

2. Human readable formatted output with optional units.
Variables that use one of the two new types can printed in human readable format using the %[hH] format character, optionally with a unit string. Lower case 'h' is used to format using decimal SI suffixes and upper case 'H" is used to format using power of two SI suffixes. For example:

		var size Int64 = Int64{3 * 1024 * 1024 * 1024, "B"}
		var speed Float64 = Float64{2100000, "bps"}
		fmt.Printf("size=%H, speed=%0.2h\n", size, speed)
		// yields:
		
		size=3 GiB, speed=2.10 Mbps

All standard SI unit prefixes are supported. In keeping with the SI standard "Ki" is used to multiply by 1024 and not "k" which is used to multiply by 1000.

This package "extends" the package flags by adding two new types and corresponding Set(string) methods.


##The Two New Types
		type Int64 struct {
			V int64
			U string
		}
		type Float64 struct {
			V float64
			U string
		}
V is for "value" and U is for the optional "units" string.

##Usage
	var blocksize hrff.Int64 = Int64{V: 4096, U: ""}

	func main() {
		flag.Var(&blocksize, "blocksize", "blocksize for copy")
		flag.Parse()
		fmt.Printf(blocksize=%H\n", blocksize)
		copy(src, dot, blocksize.V)
	}

##Bugs
1. Input arguments don't allow or record units. You have to say bs=1Ki and you can't say bs=1 KiB. Might be fixed in the future.
2. Input arguments don't allow a space between the number and the units prefix. Input format is not symmetric with respect to output format. Might be fixed in the future.
3. The are a few conflicts with hex numbers. Is "0x1d" 30 or 1 deci? Similar conflicts occur with exa, deka, femto, and atto. I suspect hex numbers aren't very useful with multiplicative suffixes so the conflict is now resolved if favor of the suffix. I think that's acceptable. It could be fixed by allowing or even forcing a space between the number and the suffix. I'll take feedback on that issue. I'm considering optionally allowing the space.


