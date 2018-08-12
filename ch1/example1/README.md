# Implement `ls` using standard library functions

The first example in APUE shows the use of the C standard library to
implement a very simple `ls`. The C standard library abstracts away a
number of issues, namely :


* the system calls themselves are unbuffered, so the management of
  providing a buffer and then parsing the directory entries from that
  buffer is a complication abstracted by the `C` standard library.
  
* different operating systems implement the `C` struct representing a
  directory entry.

* the standard library in `C` abstracts this and optimises the manner
  in which directory entries are parsed for each operating system;
  similarly the Go authors have written a standard library for Go. The
  standard library in Go; although it as to use the system calls; does
  not use the `C` standard library.

As a result there are two examples in Go for this example:


* An example using `CGo` to call the `C` standard library to implement
  a very simple `ls`.

* An example using `Go` to call the `Go` standard library to implement
  a very simple `ls`.
