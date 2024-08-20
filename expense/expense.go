package expense

import (
	"encoding/json"
	"fmt"
	"time"
)

type Date struct {
	time.Time
}

func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (d *Date) UnmarshalJSON(data []byte) (err error) {
	var s string

	if err = json.Unmarshal(data, &s); err != nil {
		return err
	}

	date, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	d.Time = date
	return nil
}

func (d Date) String() string {
	return d.Time.Format("2006-01-02")
}

func (d Date) GoString() string {
	return d.String()
}

type Expense struct {
	ID          int     `json:"id"`
	Date        Date    `json:"date"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
}

func New(id int, description string, amount float64) *Expense {
	return &Expense{
		ID:          id,
		Date:        Date{time.Now()},
		Description: description,
		Amount:      amount,
	}
}

func (e *Expense) String() string {
	if e == nil {
		return "<nil>"
	}

	return fmt.Sprintf("Expense(id=%d date=%v description=%v amount=%v)", e.ID, e.Date, e.Description, e.Amount)
}
