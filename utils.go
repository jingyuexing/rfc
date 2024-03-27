package rfc

import (
	// "encoding/hex"
	"math/rand"
	"strings"
)

const (
	hextable = "0123456789abcdef"
)

type Buffer []byte

func NewBuffer() Buffer {
	return make([]byte,0)
}

func (b Buffer) Length() int{
	return len(b)
}

func (b Buffer) String() string {
	str := make([]byte,b.Length() << 1)
	j := 0
	for _,item := range b {
		str[j] = hextable[item>>4]
		str[j+1] = hextable[item&0x0f]
		j += 2
	}
	return string(str)
}

func (b Buffer) Load(data string) Buffer{
	buf := NewBuffer()
	for i:= 0;i < len(data);i+=2 {
		cache := 0x00
		cache = strings.Index(hextable,string(data[i])) << 4
		cache = cache | strings.Index(hextable,string(data[i+1]))
		buf = append(buf,  byte(cache))
	}
	return buf
}

func (b Buffer) Rand(length int) Buffer {
	return RandBuffer(length)
}

func RandBuffer(length int) []byte {
	result := make([]byte,0)
	for i := 0;i < length;{
		temp := rand.Int()
		for temp != 0 && i < length {
			result = append(result, byte(temp & 0xff))
			temp = temp >> 8
			i++
		}
	}
	return result
}


func boolToInt(val bool) uint8 {
	if val {
		return 1
	}
	return 0
}

func Fill(target []byte,val uint) []byte {
	for val != 0 {
		c := val & 0xff
		target = append(target,byte(c))
		val = val >> 8
	}
	return target
}