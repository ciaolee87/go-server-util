package bizFormatter

import (
	"fmt"
	"time"
)

type BizTime struct {
	time.Time
}

func NewTime() *BizTime {
	return &BizTime{time.Now()}
}

func (t *BizTime) FormatISO8601() string {
	return fmt.Sprintf("%02d-%02d-%02dT%02d:%02d:%02dZ", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}

// BizTime -> json
func (t *BizTime) MarshalJSON() ([]byte, error) {
	return []byte(t.FormatISO8601()), nil
}

// json -> BizTime
func (t *BizTime) UnmarshalJSON(body []byte) error {
	if date, err := time.Parse(time.RFC3339, string(body)); err != nil {
		return err
	} else {
		t = &BizTime{date}
		return nil
	}
}
