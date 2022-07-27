package utils

import (
	"fmt"
	"testing"
)

func TestErr2Hashcode(t *testing.T) {
	n, s := Err2Hashcode(fmt.Errorf("invalide 参数"))
	fmt.Println(n, s)
}
