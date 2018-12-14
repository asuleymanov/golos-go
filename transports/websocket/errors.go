package websocket

import "github.com/pkg/errors"

//ErrClosing error returned about a closed channel
var ErrClosing = errors.New("closing")
