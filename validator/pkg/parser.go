package validator

import (
	"bytes"
	"errors"
)

func ParseCMS(data []byte) ([]byte, error) {
	if !bytes.HasPrefix(data, []byte{0x30}) {
		return nil, errors.New("invalid CMS")
	}
	return data, nil
}
