package service

import (
	"github.com/robertkrimen/otto"
)

func NewJS() *otto.Otto {
	return otto.New()
}
