package usecase

import "context"

type HelloWorld interface {
	Hello(ctx context.Context) (string, error)
}
