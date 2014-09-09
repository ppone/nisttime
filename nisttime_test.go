package nisttime

import (
	"fmt"
	"testing"
)

func TestTimeNowUTC(t *testing.T) {
	tm, s, err := TimeNowUTC()
	if err != nil {
		t.Error(err)
	}

	if s == "" {
		t.Error("Server returned no data")
	}

	fmt.Println("time is => ", tm)
}
