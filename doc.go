// Package textbuffer provides a way to buffer text data based on a number of
// writes.
//
// Sometimes there is a need to use a buffer with predictable split criteria
// so that data stays logically solid even if only a single chunk gets its way
// on the final target writer. Splitting text buffers by size is opposed to
// such behavior as text lines cut an unpredictable way.
//
// An example of such need is maintaining a log lines buffer before writing
// into in a rotated file every number of writes.
//
// This package is safe for parallel execution.
package textbuffer
