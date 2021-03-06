package vhd

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strings"
	"unicode/utf16"
	"unicode/utf8"
	"crypto/rand"
)

// https://groups.google.com/forum/#!msg/golang-nuts/d0nF_k4dSx4/rPGgfXv6QCoJ
func uuidgen() string {
	b := uuidgenBytes()
	return fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

func fmtField(name, value string) {
	fmt.Printf("%-25s%s\n", name+":", value)
}

func uuidgenBytes() []byte {
	b := make([]byte, 16)
	rand.Read(b)
	return b
}

func hexs(a []byte) string {
	return "0x" + hex.EncodeToString(a[:])
}

func uuid(a []byte) string {
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%04x",
		a[:4],
		a[4:6],
		a[6:8],
		a[8:10],
		a[10:16])
}

func uuidToBytes(uuid string) (out []byte, err error )  {
	s := strings.Replace(uuid, "-", "", -1)
	return  hex.DecodeString(s)

}

/*
	utf16BytesToString converts UTF-16 encoded bytes, in big or
 	little endian byte order, to a UTF-8 encoded string.
 	http://stackoverflow.com/a/15794113
*/
func utf16BytesToString(b []byte, o binary.ByteOrder) string {
	utf := make([]uint16, (len(b)+(2-1))/2)
	for i := 0; i+(2-1) < len(b); i += 2 {
		utf[i/2] = o.Uint16(b[i:])
	}
	if len(b)/2 < len(utf) {
		utf[len(utf)-1] = utf8.RuneError
	}
	return string(utf16.Decode(utf))
}
