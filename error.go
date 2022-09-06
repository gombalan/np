package np

import (
	"fmt"
	"runtime"
	"strings"
)

const (
	ErrZeroValue        string = "value must be non-zero"
	ErrNegativeValue    string = "value must be positive"
	ErrInvalidParameter string = "one or more parameters are invalid"
	ErrEmptyArray       string = "array is empty"
	ErrSizeNotMatch     string = "array size is not match"
)

func newError(message string, args ...arg) error {
	return create(nil, message, args)
}

// func propagate(cause error, message string) error {
// 	if cause == nil {
// 		return nil
// 	}
// 	return create(cause, message)
// }

type npError struct {
	message  string
	cause    error
	file     string
	function string
	line     int
}

type arg struct {
	key   string
	value interface{}
}

func create(cause error, message string, args []arg) error {
	err := &npError{
		message: fmt.Sprintf("%s: %v", message, args),
		cause:   cause,
	}

	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		return err
	}

	err.file, err.line = file, line

	f := runtime.FuncForPC(pc)
	if f == nil {
		return err
	}
	err.function = shortFuncName(f)

	return err
}

func shortFuncName(f *runtime.Func) string {
	longName := f.Name()

	withoutPath := longName[strings.LastIndex(longName, "/")+1:]
	withoutPackage := withoutPath[strings.Index(withoutPath, ".")+1:]

	shortName := withoutPackage
	shortName = strings.Replace(shortName, "(", "", 1)
	shortName = strings.Replace(shortName, "*", "", 1)
	shortName = strings.Replace(shortName, ")", "", 1)

	return shortName
}

func (err *npError) Error() string {
	return fmt.Sprintf("%v", *err)
}
