package main

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	sum := calculateSum(10, 11)
	assert.Equal(t, 21, sum)
}

func TestMain(t *testing.T) {
	output := captureOutput(func() {
		main()
	})
	assert.Equal(t, "Total sum=25", output)
	assert.Equal(t, "Total sum=25", output)
}

func captureOutput(f func()) string {
	orig := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f()
	os.Stdout = orig
	w.Close()
	out, _ := io.ReadAll(r)
	return string(out)
}
