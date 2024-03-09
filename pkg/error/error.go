package errors

import (
	"fmt"
	"strings"
)

func New(errType string, err error) error {
	err = fmt.Errorf("%s | %w", errType, err)
	return err
}

func ExtractError(err error) (string, error) {
	extErr := strings.Split(err.Error(), " | ")
	switch len(extErr) {
	case 2:
		return extErr[0], fmt.Errorf(extErr[1])
	default:
		return "", err
	}
}
