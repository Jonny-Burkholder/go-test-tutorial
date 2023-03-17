package example

import (
	"bytes"
	"errors"
	"time"
)

const (
	Message = "working as expected"
	Foo     = "foo"
)

var (
	ErrThisFailed = errors.New("Oh no! This failed")
	ErrUserError  = errors.New("that's not my fault, that's a user error")
)

func ExportedFunction(b *bool) (string, error) {
	if b == nil {
		return "", ErrUserError
	}
	if !*b {
		return "", ErrThisFailed
	}
	return Message, nil
}

func unexportedFunction(b bool) (string, error) {
	if !b {
		return "", ErrThisFailed
	}
	return Message, nil
}

func IsFoo(word string) bool {
	if word != Foo {
		return false
	}
	return true
}

func IsFoo2(word string) bool {
	fooBytes := []byte(Foo)
	wordBytes := []byte(word)

	if bytes.Compare(fooBytes, wordBytes) != 0 {
		return false
	}

	return true

}

func IsFoo3(word string) bool {
	time.Sleep(time.Nanosecond)
	if word != Foo {
		return false
	}
	return true
}
