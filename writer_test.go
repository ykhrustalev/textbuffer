package textbuffer_test

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/ykhrustalev/textbuffer"
	"testing"
)

func TestWriter(t *testing.T) {
	var buf bytes.Buffer

	w := textbuffer.NewWriter(&buf, 3)

	n, err := w.Write([]byte("foo"))
	assert.NoError(t, err)
	assert.Equal(t, 3, n)
	assert.Equal(t, "", buf.String())

	n, err = w.WriteString("hello")
	assert.NoError(t, err)
	assert.Equal(t, 5, n)
	assert.Equal(t, "", buf.String())

	n, err = w.WriteString("")
	assert.NoError(t, err)
	assert.Equal(t, 0, n)
	assert.Equal(t, "foohello", buf.String())

	n, err = w.WriteString("1")
	assert.NoError(t, err)
	assert.Equal(t, 1, n)
	assert.Equal(t, "foohello", buf.String())

	n, err = w.WriteString("2")
	assert.NoError(t, err)
	assert.Equal(t, 1, n)
	assert.Equal(t, "foohello", buf.String())

	n, err = w.WriteString("3")
	assert.NoError(t, err)
	assert.Equal(t, 1, n)
	assert.Equal(t, "foohello123", buf.String())

	n, err = w.WriteString("bar")
	assert.NoError(t, err)
	assert.Equal(t, 3, n)

	assert.NoError(t, w.Flush())
	assert.Equal(t, "foohello123bar", buf.String())
}

func Example_basic() {
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
}
