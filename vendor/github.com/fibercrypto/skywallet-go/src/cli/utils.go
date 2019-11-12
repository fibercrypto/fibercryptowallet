package cli

import (
	"errors"
)

func parseBool(s string) (*bool, error) {
	var b bool
	switch s {
	case "true":
		b = true
	case "false":
		b = false
	case "":
		return nil, nil
	default:
		return nil, errors.New("Invalid boolean argument")
	}
	return &b, nil
}
