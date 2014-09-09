package nisttime

import (
	"fmt"
	"testing"
)

func TestTimeNowUTC(t *testing.T) {
	time, err := TimeNowUTC()
	if err != nil {
		t.Error(err)
	}

	if time == "" {
		t.Error("Server returned no data")
	}

	fmt.Println("time is => ", time)
}
