package protowire

import (
	"github.com/ixbasANT/gord/app/appmessage"
	"github.com/pkg/errors"
)

func (x *RPCError) toAppMessage() (*appmessage.RPCError, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "RPCError is nil")
	}
	return &appmessage.RPCError{Message: x.Message}, nil
}
