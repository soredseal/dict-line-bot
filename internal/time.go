package internal

import (
	"strconv"
	"time"
)

type Time struct {
	time.Time
}


func (t *Time) UnmarshalJSON(b []byte) error {
	timestamp := []rune(string(b))[0:10]
	epoch, err := strconv.ParseInt(string(timestamp), 10, 64)
	if err != nil {
		return err
	}
	*t = Time{time.Unix(epoch, 0)}
	return nil
}

