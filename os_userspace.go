//go:build !linux && !openbsd && !windows && !freebsd
// +build !linux,!openbsd,!windows,!freebsd

package wgctrl

import (
	"github.com/Biezax/wgctrl/internal/wginternal"
	"github.com/Biezax/wgctrl/internal/wguser"
	"github.com/Biezax/wgctrl/wgtypes"
)

// newClients configures wginternal.Clients for systems which only support
// userspace WireGuard implementations.
func newClients(clientType wgtypes.ClientType) ([]wginternal.Client, error) {
	c, err := wguser.New(clientType)
	if err != nil {
		return nil, err
	}

	return []wginternal.Client{c}, nil
}
