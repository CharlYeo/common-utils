package string_util

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	resp := Hello()
	fmt.Println("resp is :", resp)
}
