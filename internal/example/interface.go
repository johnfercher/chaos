package example

import "context"

type Example interface {
	Add(ctx context.Context, id string) error
}
