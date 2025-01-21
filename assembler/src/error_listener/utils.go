package error_listener

import "errors"

func WrapErrorMessages(errs ...string) error {
	errLen := len(errs)
	if errLen == 0 {
		return nil
	}
	s := "Assemble failed.\r\n"
	s += "Syntax Error(s):\r\n"
	for i, e := range errs {
		s += "\t" + e + "\r"
		if i != errLen-1 {
			s += "\n"
		}
	}
	return errors.New(s)
}

func WrapErrors(errs ...error) error {
	errLen := len(errs)
	if errLen == 0 {
		return nil
	}
	s := "Assemble failed.\r\n"
	s += "Syntax Error(s):\r\n"
	for i, e := range errs {
		s += "\t" + e.Error() + "\r"
		if i != errLen-1 {
			s += "\n"
		}
	}
	return errors.New(s)
}
