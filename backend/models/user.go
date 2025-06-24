package models

import "fmt"

func (u *User) Validate() error {
	if u.PhoneNumber == "" {
		return fmt.Errorf("phone number is required")
	}
	return nil
}
