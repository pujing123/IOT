package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"log"
	"math"
	"strconv"
)

func main(){
	src := strconv.FormatInt(16, 16)
	test, er := hex.DecodeString(src)
	log.Println(test, er)

	bytesBuffer := bytes.NewBuffer([]byte{})
	err := binary.Write(bytesBuffer, binary.BigEndian, int32(100))
	b:= bytesBuffer.Bytes()
	log.Println(b,err)
	log.Println(IntToBytes(15))


	s := "C40C5253"
	log.Println(parse_read(s))
	log.Println(parse_math(s))

}

func IntToBytes(n int) []byte {
	data := int64(n)
	bytebuf := bytes.NewBuffer([]byte{})
	binary.Write(bytebuf, binary.BigEndian, data)
	return bytebuf.Bytes()
}

func parse_read(s string) (f float32, err error) {
	b, err := hex.DecodeString(s)
	if err != nil {
		return
	}
	buf := bytes.NewReader(b)
	err = binary.Read(buf, binary.BigEndian, &f)
	return
}

func parse_math(s string) (f float32, err error) {
	i, err := strconv.ParseUint(s, 16, 32)
	if err != nil {
		return
	}
	f = math.Float32frombits(uint32(i))
	return
}

