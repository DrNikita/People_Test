package status

import (
	"errors"
	"fmt"
)

func PeopleNotFound() error {
	return errors.New("people not found")
}

func NonExistentId(id int) error {
	return errors.New(fmt.Sprintf("id : %d does not exist", id))
}
