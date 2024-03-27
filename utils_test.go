package rfc_test

import (
	"fmt"
	"testing"

	"github.com/jingyuexing/rfc"
)

func TestRandBuffer(t *testing.T) {
	buf := rfc.NewBuffer()
	buf = buf.Rand(128)
	str := buf.String()
	buf2 := rfc.NewBuffer()
	buf2 = buf2.Load(str)
	if buf.String() != buf2.String() {
		t.Error(fmt.Sprintf("the buffer1 [%s] should same with buffer2[%s]",buf,buf2))
	}
}