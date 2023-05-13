//go:build js
// +build js

package console

func checkConsole() (Console, error) {
	return nil, "unsupported"
}

func newMaster(f File) (Console, error) {
	return nil, "unsupported"
}
