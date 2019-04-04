package errors

import "fmt"

type cause interface {
	Cause() error
}

type errorApp struct {
	msg    string
	caller string
	cause  error
}

func (e errorApp) Error() string {
	return fmt.Sprintf("%s : %s - Caused by : %s", e.caller, e.msg, e.cause.Error())
}

func (e errorApp) Cause() error {
	return e.cause
}

func New(caller string, msg string, cause error) error {
	return &errorApp{msg: msg, caller: caller, cause: cause}
}

func HandleError(f func(args ...interface{}), err error) error {
	f(err)
	return err
}

func Cause(err error) error {
	if err == nil {
		return nil
	}

	c, ok := err.(cause)
	if !ok {
		return nil
	}

	return c.Cause()
}
