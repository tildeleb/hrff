#siflags

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


