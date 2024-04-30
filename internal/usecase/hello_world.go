package usecase

import "context"

type HelloWorldUC struct {
}

func NewHelloWorld() *HelloWorldUC {
	return &HelloWorldUC{}
}

func (hw *HelloWorldUC) Hello(_ context.Context) (string, error) {
	return "Hello, world!", nil
}
