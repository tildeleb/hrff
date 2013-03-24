siflags
=======

Allows SI unit prefixes to be used as suffixes to numeric flags in Go Lang CLI flags by adding two new types and methods.
For example you can use flags like the blocksize argument below:

	% copy file1 file2 blocksize=1Ki