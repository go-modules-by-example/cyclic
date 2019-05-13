package b_test

import (
	"fmt"
	"github.com/go-modules-by-example/cyclic/a"
	"testing"
)

func TestUsingA(t *testing.T) {
	fmt.Printf("Here is A: %v\n", a.AName)
}
