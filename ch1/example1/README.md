# Implement `ls` Using Standard Library Functions

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
  standard library in Go; although it has to use the `C` system calls;
  does not use the `C` standard library.

## Code Examples 

As a result there are two examples in Go for this example:


* An example using `CGo` to call the `C` standard library to implement
  a very simple `ls`.

* An example using `Go` to call the `Go` standard library to implement
  a very simple `ls`.


## Functional Differences

* The C standard library call's directory entries are not returned in a determined manner. They are not ordered and may be returned in any order. The underlying unbuffered system call also does not return the entries in a determined or sorted order.

* The Go standard library returns the directory entries in a sorted lexical order. As the underlying system call does not sort the entries, this sorting is done within the standard library after calling the underlying unbuffered system call.

* The C standard library returns all directory entries. This includes the directory entries for `.` and `..`. The underlying system call also returns these entries.

* The Go standard library does not return `.` and `..` in the list of directory entries. This means that the Go standard library removes these entries after calling the underlying system call, as the system call does return these entries.

