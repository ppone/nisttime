package nisttime

import (
	"fmt"
	"net"
	"strings"
	"time"
)

const NIST_TIME_URL = "time.nist.gov:13"
const BYTESIZE = 52

func TimeNowUTC() (time.Time, string, error) {

	conn, err := net.Dial("tcp", NIST_TIME_URL)
	if err != nil {
		fmt.Println("ERROR while connecting to " + NIST_TIME_URL + " => " + err.Error())
		return time.Time{}, "", err

	}
	b := make([]byte, BYTESIZE)

	n, err := conn.Read(b)

	if err != nil {
		fmt.Println("ERROR while reading data =>", err)
		return time.Time{}, "", err
	}

	temp := strings.TrimSpace(string(b[0:n])[6:25])
	format := "06-01-02 15:04:05"

	s, err := time.Parse(format, temp)

	if err != nil {
		return time.Time{}, "", err
	}

	return s, temp, nil

}
