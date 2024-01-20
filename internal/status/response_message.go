package status

import "errors"

func DbError() error {
	return errors.New("db connection error")
}
func PeopleNotFound() error {
	return errors.New("people not found")
}
