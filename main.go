package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

//ENCODER STARTS
//  CMD    NO_OF_PARAM   DETAILS
//  PUT    2             KEY VALUE
//  GET    1             KEY
//  DEL    1             KEY
//ENCODER ENDS

//  PARSER START
//
//  PARSER END

//  ERROR CODES START --Will be coming as response
//    ERR_001 INVALID_CMD
//    ERR_OO2 KEY_NOT_FOUND
//    ERR_999 ERR_UNKNOWN
//  ERROR CODES END

func main() {
	fmt.Println("Hello World!")

	// create a tcp connectioncons
	conn, err := net.Dial("tcp", "localhost:6379")
	defer conn.Close()
	if err != nil {
		panic(err)
	}

	tlv := TLV{
		Tag:    uint16(1),
		Length: uint16(len("Hello")),
		Value:  []byte("Hello"),
	}

	_, err = conn.Write(tlv.Encode())
	if err != nil {
		panic(nil)
	}
}

type TLV struct {
	Tag    uint16 //2bit
	Length uint16 //bit
	Value  []byte //as per length
}

func (t *TLV) Encode() []byte {
	buf := make([]byte, 4+len(t.Value))
	binary.BigEndian.PutUint16(buf, t.Tag)
	binary.BigEndian.PutUint16(buf[2:], t.Length)
	copy(buf[4:], t.Value)
	fmt.Println(buf)
	return buf
}
