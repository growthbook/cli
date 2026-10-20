// Hand-written (NOT generated). End-to-end tests for confirmation holds.
package cli

import (
	"net/http"
	"strings"
	"testing"
)

const heldBody = `{"message":"Held for confirmation","confirmation":{"id":"cfm_1","status":"pending","actions":[{"action":"project.delete","environments":[]}],"summary":"Delete project Alpha","url":"https://app.growthbook.io/confirmations/cfm_1","expiresAt":"2026-10-21T00:00:00Z","response":null,"note":null}}`

func respond(status int, body string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_, _ = w.Write([]byte(body))
	}
}

func TestE2E_HeldWrite_ExitsPendingWithNotice(t *testing.T) {
	// json is the raw passthrough; yaml goes through the typed SDK response.
	for _, format := range []string{"json", "yaml"} {
		t.Run(format, func(t *testing.T) {
			out, _, err := runCLI(t, respond(http.StatusAccepted, heldBody),
				"projects", "delete", "--id", "prj_1", "--output-format", format)

			if code := ExitCode(err); code != ExitConfirmationPending {
				t.Fatalf("exit code = %d, want %d (err: %v)", code, ExitConfirmationPending, err)
			}
			for _, want := range []string{
				"Delete project Alpha",
				"https://app.growthbook.io/confirmations/cfm_1",
				"meta get-confirmation --id cfm_1",
			} {
				if !strings.Contains(err.Error(), want) {
					t.Errorf("notice missing %q:\n%s", want, err)
				}
			}
			if !strings.Contains(out, "cfm_1") {
				t.Errorf("stdout should still carry the confirmation body, got:\n%s", out)
			}
		})
	}
}

func TestE2E_GetConfirmation_ExitCodes(t *testing.T) {
	cases := []struct {
		name, confirmation string
		code               int
		notice             string
	}{
		{"pending", `"status":"pending","response":null,"note":null`, ExitConfirmationPending, "Held for human confirmation"},
		{"running", `"status":"running","response":null,"note":null`, ExitConfirmationPending, "still running"},
		{"completed", `"status":"completed","response":{"status":200,"body":{}},"note":null`, 0, ""},
		{"completed but failed", `"status":"completed","response":{"status":409,"body":{}},"note":null`, 1, "failed with HTTP 409"},
		{"rejected", `"status":"rejected","response":null,"note":"Not during the freeze"`, 1, "Not during the freeze"},
		{"expired", `"status":"expired","response":null,"note":null`, 1, "Expired"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			body := `{"confirmation":{"id":"cfm_1","actions":[],"summary":"Delete project Alpha","url":"https://app.growthbook.io/confirmations/cfm_1","expiresAt":"2026-10-21T00:00:00Z",` + tc.confirmation + `}}`
			_, _, err := runCLI(t, respond(http.StatusOK, body),
				"meta", "get-confirmation", "--id", "cfm_1", "--output-format", "json")

			if tc.code == 0 {
				if err != nil {
					t.Fatalf("want success, got: %v", err)
				}
				return
			}
			if code := ExitCode(err); code != tc.code {
				t.Fatalf("exit code = %d, want %d (err: %v)", code, tc.code, err)
			}
			if !strings.Contains(err.Error(), tc.notice) {
				t.Errorf("notice missing %q:\n%s", tc.notice, err)
			}
		})
	}
}

func TestE2E_OrdinaryWrite_Unaffected(t *testing.T) {
	_, _, err := runCLI(t, respond(http.StatusOK, `{"deletedId":"prj_1"}`),
		"projects", "delete", "--id", "prj_1", "--output-format", "json")
	if err != nil {
		t.Fatalf("plain success should exit 0, got: %v", err)
	}
}
