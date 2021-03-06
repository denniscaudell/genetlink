//+build linux

package genltest_test

import (
	"os"
	"syscall"
	"testing"

	"github.com/mdlayher/genetlink"
	"github.com/mdlayher/genetlink/genltest"
	"github.com/mdlayher/netlink"
)

func TestConnLinuxReceiveError(t *testing.T) {
	c := genltest.Dial(func(_ genetlink.Message, _ netlink.Message) ([]genetlink.Message, error) {
		return nil, genltest.Error(int(syscall.EPERM))
	})
	defer c.Close()

	_, _, err := c.Receive()
	if !os.IsPermission(err) {
		t.Fatalf("expected permission denied error, but got: %v", err)
	}
}
