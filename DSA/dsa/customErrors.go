package dsa

import "fmt"

type DataNotFoundError struct {
	arg     int
	message string
}

// to be considered as a error a struct has to implement errors interface i.e implement Error() function
// errors.As() ??
func (derr *DataNotFoundError) Error() string {
	return fmt.Sprintf("%d - %s", derr.arg, derr.message)
}
