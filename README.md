#siflags

***WARNING: This project under heavy development right now with new capabilities on the way. The package name will be changing as will some of the datastructures. Use at your own risk as the next version will be incompatible in some significant ways. -- leb***

Allows SI unit prefixes to be used as mutiplicative suffixes to numeric flags in Go Lang. This package extends the package "flags" by adding two new types and cooresponding Set(string) methods.
For example you can use flags like the blocksize argument below:

	% # set block size to 1024
	% copy file1 file2 blocksize=1Ki

All standard SI prefixes are supported. In keeping with the SI standard "Ki" is used to multiply by 1024 and not "k" which is used to multiply by 1000.

##The Two New Types
	Int64Value
	Float64Value

##Usage
	var _blocksize siflags.Int64Value = 512
	var blocksize int64

	func main() {
		flag.Var(&_blocksize, "blocksize", "blocksize for copy")
		flag.Parse()
		blocksize = int64(_blocksize)
	}

##Bugs
The are a few conlicts with hex numbers. Is "0x1d" 30 or 1 deci? Similar conflicts occur with exa, deka, femto, and atto. I suspect hex numbers aren't very useful with mutiplicative suffixes so the conflict is now resolved if favor of the suffix. I think that's acceptable. It could be fixed by allowing or even forcing a space between the number and the suffix. I'll take feedback on that issue. I'm considering optionally allowing the space.


