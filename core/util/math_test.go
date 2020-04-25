package util_test

import (
	"github.com/RuchDB/chaos/util"
	"testing"
)

func TestIntAdd(t *testing.T) {
	if sum := util.IntAdd(1, 1); sum != 2 {
		t.Errorf("IntAdd(1, 1): 2 expected, but %d got", sum)
	}
}
