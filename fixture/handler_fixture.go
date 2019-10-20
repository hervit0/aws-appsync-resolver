package fixture

import (
	"context"
	"errors"
)

type Handler struct {
	Error string
}

func (handler Handler) Invoke(ctx context.Context, payload RequestFixture) ([]byte, error) {
	return nil, errors.New(handler.Error)
}
