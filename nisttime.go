package nisttime

import (
	"fmt"
	"net"
)

const NIST_TIME_URL = "time.nist.gov:13"
const BYTESIZE = 52

func TimeNowUTC() (string, error) {

	conn, err := net.Dial("tcp", NIST_TIME_URL)
	if err != nil {
		fmt.Println("ERROR while connecting to " + NIST_TIME_URL + " => " + err.Error())
		return "", nil

	}
	b := make([]byte, BYTESIZE)

	n, err := conn.Read(b)

	if err != nil {
		fmt.Println("ERROR while reading data =>", err)
		return "", err
	}

	return string(b[0:n]), nil

}
