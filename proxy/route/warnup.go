package route

import (
	"overlord/proxy/proto"
)

// WarnUp cold cache warn up forward.
type WarnUp struct {
	cold proto.Forwarder
	warn proto.Forwarder
}
