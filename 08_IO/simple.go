package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var b bytes.Buffer
	b.Write([]byte("hello"))

	fmt.Fprintf(&b, "world")

	b.WriteTo(os.Stdout)
}
