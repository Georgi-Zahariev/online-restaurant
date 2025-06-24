package models

import (
	"fmt"
)

func (o *Order) Validate() error {
	if o.Price <= 0 {
		return fmt.Errorf("price must be greater than zero")
	}
	if o.Status == "" {
		return fmt.Errorf("status is required")
	}
	if o.UserID == "" {
		return fmt.Errorf("user ID is required")
	}
	if o.DayAndTime.IsZero() {
		return fmt.Errorf("day and time is required")
	}
	return nil
}
