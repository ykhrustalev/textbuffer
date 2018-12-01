# Textbuffer 
[![GoDoc](https://godoc.org/github.com/ykhrustalev/textbuffer?status.svg)](https://godoc.org/github.com/ykhrustalev/textbuffer)
[![Go Report Card](https://goreportcard.com/badge/github.com/ykhrustalev/textbuffer)](https://goreportcard.com/report/github.com/ykhrustalev/textbuffer)

Package textbuffer provides a way to buffer text data based on a number of
writes.

Sometimes there is a need to use a buffer with predictable split criteria
so that data stays logically solid even if only a single chunk gets its way
on the final target writer. Splitting text buffers by size is opposed to
such behavior as text lines cut an unpredictable way.

An example of such need is maintaining a log lines buffer before writing
into in a rotated file every number of writes.

Example:
```go
var buf bytes.Buffer

w := textbuffer.NewWriter(&buf, 2)

_, _ = w.Write([]byte("first\n"))
fmt.Println("Action#1")
fmt.Println(buf.String())

_, _ = w.WriteString("second\n")
fmt.Println("Action#2")
fmt.Println(buf.String())

_, _ = w.Write([]byte("third\n"))
fmt.Println("Action#3")
fmt.Println(buf.String())
// Output:
// Action#1
//
// Action#2
// first
// second
//
// Action#3
// first
// second
```
