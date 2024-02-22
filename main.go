package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Hello World!")

	// create a tcp connectioncons
	conn, err := net.Dial("tcp", "localhost:6379")
	defer conn.Close()
	if err != nil {
		panic(err)
	}

	//ENCODER STARTS
	// CMD    NO_OF_PARAM   DETAILS
	// PUT    2             KEY VALUE
	// GET    1             KEY
	// DEL    1             KEY
	//ENCODER ENDS

	//  PARSER START
	//
	//  PARSER END

	//  ERROR CODES START --Will be coming as response
	//  ERR_001 INVALID_CMD
	//  ERR_OO2 KEY_NOT_FOUND
	//  ERR_999 ERR_UNKNOWN
	//  ERROR CODES END
	data := []byte("PUT\r\nKEY\r\nVALUE\r\n")
	_, err = conn.Write(data)
	if err != nil {
		panic(nil)
	}
}
