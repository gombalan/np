package np

import (
	"fmt"
	"runtime"
	"strings"
)

const (
	ErrZeroValue           string = "value is zero"
	ErrNegativeValue       string = "value is negative"
	ErrZeroOrNegativeValue string = "value is zero or negative"
	ErrInvalidParameter    string = "one or more parameters are invalid"
	ErrEmptyArray          string = "array is empty"
	ErrSizeNotMatch        string = "array size is not matched"
)

func newError(errCode string, errDescription string) error {
	return createError(nil, errCode+": "+errDescription)
}

func propagateError(cause error, errDescription string) error {
	if cause == nil {
		return nil
	}
	return createError(cause, errDescription)
}

type npError struct {
	description string
	cause       error
	file        string
	function    string
	line        int
}

func createError(cause error, description string) error {
	err := &npError{
		description: description,
		cause:       cause,
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
