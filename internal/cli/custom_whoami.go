// Hand-written (NOT generated). Helper for the patched whoami command.
package cli

import (
	"strconv"

	"github.com/growthbook/cli/internal/flagutil"
	"github.com/growthbook/cli/internal/sdk"
	"github.com/spf13/cobra"
)

// resolvedServerURL mirrors client.NewClient's server resolution so whoami can
// report the server requests will actually go to: --server-url wins, then a
// valid --server index into sdk.ServerList, otherwise the default (ServerList[0]).
func resolvedServerURL(cmd *cobra.Command) string {
	if u, _ := flagutil.GetStringFlag(cmd, "server-url"); u != "" {
		return u
	}
	if s, _ := flagutil.GetStringFlag(cmd, "server"); s != "" {
		if idx, err := strconv.Atoi(s); err == nil && idx >= 0 && idx < len(sdk.ServerList) {
			return sdk.ServerList[idx]
		}
	}
	return sdk.ServerList[0]
}
