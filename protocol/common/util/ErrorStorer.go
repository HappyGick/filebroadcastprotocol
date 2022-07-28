package util

type ErrorStorer interface {
	GetError() error
}
