package main

import (
	"fmt"
	"github.com/golang/example/stringutil"
	"testing"
)

func TestSum(t *testing.T) {
	total := Sum(5, 9)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}

func ExampleReverse() {
	fmt.Println(stringutil.Reverse("hello_"))
	// Output: olleh
}