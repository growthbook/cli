// Hand-written (NOT generated). Surfaces GrowthBook confirmation holds: a write
// the server holds for a person to approve (HTTP 202), and the outcome reported
// by `meta get-confirmation`. Either way the response body still prints as
// usual; this adds a stderr notice and an exit code a script or agent can't miss.
package cli

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/spf13/cobra"
)

// ExitConfirmationPending is the exit code while an action waits on a person.
const ExitConfirmationPending = 3

type confirmation struct {
	ID        string  `json:"id"`
	Status    string  `json:"status"`
	Summary   string  `json:"summary"`
	URL       string  `json:"url"`
	ExpiresAt string  `json:"expiresAt"`
	Note      *string `json:"note"`
	Response  *struct {
		Status float64 `json:"status"`
	} `json:"response"`
}

// confirmationError ends a command with a confirmation notice and its exit code.
type confirmationError struct {
	code int
	msg  string
}

func (e *confirmationError) Error() string { return e.msg }

// ExitCode maps an error returned by Execute to the process exit code.
func ExitCode(err error) int {
	var c *confirmationError
	if errors.As(err, &c) {
		return c.code
	}
	return 1
}

var (
	lastConfirmation atomic.Pointer[confirmation]
	installTransport sync.Once
)

// registerConfirmationHandling records confirmations seen on the wire and
// reports them once the command has printed its output.
func registerConfirmationHandling(rootCmd *cobra.Command) {
	// The generated SDK client leaves Transport nil, so it uses http.DefaultTransport.
	installTransport.Do(func() {
		http.DefaultTransport = confirmationTransport{next: http.DefaultTransport}
	})

	genPostRun := rootCmd.PersistentPostRunE
	rootCmd.PersistentPostRunE = func(cmd *cobra.Command, args []string) error {
		if genPostRun != nil {
			if err := genPostRun(cmd, args); err != nil {
				return err
			}
		}
		return confirmationOutcome(cmd.Root().Name(), lastConfirmation.Swap(nil))
	}
}

type confirmationTransport struct{ next http.RoundTripper }

func (t confirmationTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	res, err := t.next.RoundTrip(req)
	if err != nil || !mayCarryConfirmation(req, res) {
		return res, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}
	res.Body = io.NopCloser(bytes.NewReader(body))

	var env struct {
		Confirmation *confirmation `json:"confirmation"`
	}
	if json.Unmarshal(body, &env) == nil && env.Confirmation != nil && env.Confirmation.ID != "" {
		lastConfirmation.Store(env.Confirmation)
	}
	return res, nil
}

// mayCarryConfirmation matches a held write (202) or a `meta get-confirmation` read.
func mayCarryConfirmation(req *http.Request, res *http.Response) bool {
	if res.StatusCode == http.StatusAccepted {
		return true
	}
	return req.Method == http.MethodGet && res.StatusCode == http.StatusOK &&
		strings.Contains(req.URL.Path, "/v1/confirmations/")
}

func confirmationOutcome(cli string, c *confirmation) error {
	if c == nil {
		return nil
	}
	check := fmt.Sprintf("%s meta get-confirmation --id %s --wait 30", cli, c.ID)
	switch c.Status {
	case "pending":
		return &confirmationError{ExitConfirmationPending, fmt.Sprintf(
			"Held for human confirmation: %s\nA person must approve it in GrowthBook before %s: %s\nThen check the outcome: %s",
			c.Summary, c.ExpiresAt, c.URL, check)}
	case "running":
		return &confirmationError{ExitConfirmationPending, fmt.Sprintf(
			"Confirmed and still running: %s\nCheck again: %s", c.Summary, check)}
	case "rejected":
		msg := "Rejected in GrowthBook: " + c.Summary
		if c.Note != nil && *c.Note != "" {
			msg += "\nNote: " + *c.Note
		}
		return &confirmationError{1, msg}
	case "expired":
		return &confirmationError{1, "Expired before anyone confirmed it: " + c.Summary}
	case "completed":
		if c.Response != nil && c.Response.Status >= 400 {
			return &confirmationError{1, fmt.Sprintf(
				"Confirmed, but the action failed with HTTP %d: %s", int(c.Response.Status), c.Summary)}
		}
	}
	return nil
}
