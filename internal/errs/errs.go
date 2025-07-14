package errs

import (
	"fmt"
)

type Error struct {
	Code int    `json:"code"`
	Msg  string `json:"err"`
}

func (this *Error) Error() string {
	return fmt.Sprintf("code: %v, err: %s", this.Code, this.Msg)
}
