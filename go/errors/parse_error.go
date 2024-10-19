package errors

import "fmt"

type ParseError struct {
	TargetType string
	Input      string
}

func (e *ParseError) Error() string {
	return fmt.Sprintf(`Could not parse the following string as '%s':
%s`, e.TargetType, e.Input)
}
