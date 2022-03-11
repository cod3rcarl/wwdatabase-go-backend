package client

import (
	"errors"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrInternal      = errors.New("service internal error")
	ErrUnavailable   = errors.New("service unavailable")
	ErrUnknown       = errors.New("service error")
	ErrNoNameOrID    = errors.New("must provide name or ID")
	ErrDateIncorrect = errors.New("date incorrectly formatted")
)

func (c *Client) handleErr(err error) error {
	if s, ok := status.FromError(err); ok {
		if s.Code() == codes.NotFound {
			return err
		}
		if s.Code() == codes.Unavailable {
			c.logger.Error("service unavailable", zap.Error(err))

			return ErrUnavailable
		}
	}

	c.logger.Error("service error", zap.Error(err))

	return ErrUnknown
}
