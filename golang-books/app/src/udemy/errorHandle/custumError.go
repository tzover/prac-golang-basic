package errorHandle

import "fmt"

type UserNotFound struct {
	UserName string
}

func (e *UserNotFound) Error() string {
	return fmt.Sprintf("User not found : %v", e.UserName)
}

func myFunc() error {
	ok := false

	if ok {
		return nil
	}
	return &UserNotFound{UserName: "Mike"}
}
