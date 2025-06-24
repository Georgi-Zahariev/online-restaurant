package models

import "fmt"

func (d *Dish) Validate() error {
	if d.Name == "" {
		return fmt.Errorf("name is required")
	}
	if d.Price <= 0 {
		return fmt.Errorf("price must be greater than zero")
	}
	if d.CategoryID == "" {
		return fmt.Errorf("category ID is required")
	}
	return nil
}
