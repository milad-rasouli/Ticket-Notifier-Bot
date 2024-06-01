package entity

import "time"

type CustomTime struct {
	time.Time
}

const ctLayout = "2006-01-02T15:04:05"

func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := string(b)
	s = s[1 : len(s)-1] // Remove quotes
	t, err := time.Parse(ctLayout, s)
	if err != nil {
		return err
	}
	ct.Time = t
	return nil
}

type Ticket struct {
	Corporation   string     `json:"corporation"`
	DepartureTime CustomTime `json:"departureTime"`
	ArrivalTime   CustomTime `json:"arrivalTime"`
	FromCity      string     `json:"fromCity"`
	ToCity        string     `json:"toCity"`
	Price         int64      `json:"price"`
	Capacity      int32      `json:"capacity"`
	BusType       string     `json:"busType"`
}

type Bus struct {
	Buses []Ticket `json:"buses"`
}
