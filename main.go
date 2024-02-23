package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

//TODO: LATER to be made as CLI APP
//TODO: OPTIMIZE reponse length

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

	remoteAdr := conn.RemoteAddr().String()
	// keep the server open until client do exit
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(remoteAdr + "> ") //Printing the console text
		// read the respnse from the server
		// setting maximum 1024 buffer
		// take input from scanner
		line, err := reader.ReadString('\n')
		now := time.Now()
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("read line: %s-\n", line)
		// clean line variable, remove space and \n
		line = cleanLine(line)
		tlv := TLV{
			Tag:    uint16(2),
			Length: uint16(len(line)),
			Value:  []byte(line), //should be more than 3 character
		}

		_, err = conn.Write(tlv.Encode())
		if err != nil {
			panic(nil)
		}
		buf := make([]byte, 1024)
		_, err = conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				log.Panic("Server not abailable, Aborting")
			}
			log.Println(err)
		}

		// decodingin the server response
		resp, err := Decode(buf)
		if err != nil {
			log.Println(err)
		}
		if int(resp.Tag) != 0 {
			fmt.Println("ERROR, ", string(resp.Value))
		} else {
			fmt.Println("OK, ", string(resp.Value), "TIME: ", time.Now().Sub(now))
		}
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
	// fmt.Println(buf)
	return buf
}

func cleanLine(line string) string {
	// Remove leading and trailing whitespaces
	cleanedLine := strings.TrimSpace(line)
	// Remove newline characters
	cleanedLine = strings.ReplaceAll(cleanedLine, "\n", "")
	// Remove space characters
	cleanedLine = strings.ReplaceAll(cleanedLine, " ", "")

	return cleanedLine
}

func Decode(data []byte) (*TLV, error) {
	//four bit are reserved for Key, and length
	if len(data) <= 5 {
		return nil, fmt.Errorf("insufficient data")
	}
	cmd := binary.BigEndian.Uint16(data[:2])
	length := binary.BigEndian.Uint16(data[2:4])

	if len(data) < int(length)+4 {
		return nil, fmt.Errorf("insufficient data for TLV value decoding")
	}

	// fmt.Println("DATA: ", data)
	nData := data[4:]

	// fmt.Println("DATA: ", nData, " CMD: ", cmd, " LENGTH: ", length)

	tlv := TLV{
		Tag:    cmd,
		Length: length,
		Value:  nData,
	}
	return &tlv, nil
}
