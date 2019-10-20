package fixture

import (
	"context"
	"errors"
	"log"
)

type Handler struct {
	Error string
}

func (handler Handler) Handle(ctx context.Context, payload RequestFixture) (ResponseFixture, error) {
	log.Printf("Payload: %+v", payload)
	log.Printf("About to respond")
	return ResponseFixture{Animal: "yo"}, errors.New(handler.Error)
}
