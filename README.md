nisttime
========

Go program to pull the current time in UTC from the NIST server

Install and/or Update

```
go get -u github.com/ppone/nisttime

```


Example Program

```Go

package main

import "fmt"
import "github.com/ppone/nisttime"

func main() {

	t, _, err := nisttime.TimeNowUTC()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(t, t.Unix())
}

```