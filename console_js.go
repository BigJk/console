//go:build js
// +build js

package console

import (
	"errors"
)

func checkConsole(f File) error {
	return errors.New("unsupported")
}

func newMaster(f File) (Console, error) {
	return nil, errors.New("unsupported")
}
