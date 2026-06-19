package customcfg

import (
	"testing"
	"time"
)

func TestCompareSemver(t *testing.T) {
	cases := []struct {
		a, b string
		want int
	}{
		{"1.0.0", "1.0.0", 0},
		{"1.0.1", "1.0.0", 1},
		{"1.0.0", "1.0.1", -1},
		{"1.2.0", "1.10.0", -1}, // numeric, not lexical
		{"2.0.0", "1.9.9", 1},
		{"v1.0.0", "1.0.0", 0},     // leading v tolerated
		{"1.0.0", "dev", 1},        // non-numeric → 0.0.0
		{"1.0.0-rc1", "1.0.0", 0},  // pre-release suffix on a segment ignored
	}
	for _, c := range cases {
		if got := compareSemver(c.a, c.b); got != c.want {
			t.Errorf("compareSemver(%q, %q) = %d, want %d", c.a, c.b, got, c.want)
		}
	}
}

func TestShouldNudge(t *testing.T) {
	rel := func(v, d string) *cliRelease { return &cliRelease{Version: v, SpecDate: d} }
	srv := func(d string) *serverBuild { return &serverBuild{Date: d} }

	if _, ok := shouldNudge("1.0.0", "", rel("1.0.0", ""), nil); ok {
		t.Error("no nudge expected when already on the latest version")
	}
	if _, ok := shouldNudge("1.0.0", "", nil, nil); ok {
		t.Error("no nudge expected when the latest release is unknown")
	}
	if _, ok := shouldNudge("1.0.0", "", rel("1.1.0", ""), nil); !ok {
		t.Error("nudge expected when a newer version exists and there are no dates to gate on")
	}
	// Backend predates what the newer CLI targets → must NOT push the upgrade.
	if _, ok := shouldNudge("1.0.0", "2026-01-01", rel("1.1.0", "2026-06-01"), srv("2026-03-01")); ok {
		t.Error("no nudge expected when the backend is older than the newer CLI's generated-against date")
	}
	// Backend is new enough → nudge.
	if msg, ok := shouldNudge("1.0.0", "2026-01-01", rel("1.1.0", "2026-06-01"), srv("2026-07-01")); !ok {
		t.Error("nudge expected when the backend is new enough for the newer CLI")
	} else if msg == "" {
		t.Error("nudge message should be non-empty")
	}
	// Patch-only bump → no nudge (usually codegen-only, no new capability).
	if _, ok := shouldNudge("1.2.0", "", rel("1.2.1", ""), nil); ok {
		t.Error("no nudge expected for a patch-only bump")
	}
	// Major bump → nudge.
	if _, ok := shouldNudge("1.2.0", "", rel("2.0.0", ""), nil); !ok {
		t.Error("nudge expected for a major bump")
	}
}

func TestParseSpecDate(t *testing.T) {
	cases := []struct {
		body, want string
	}{
		{"Release notes\n<!-- gb-spec-date: 2026-06-18 -->\n", "2026-06-18"},
		{"gb-spec-date:2026-01-02", "2026-01-02"},   // no space, no trailing marker
		{"some notes without the marker", ""},        // absent → empty
		{"<!-- gb-spec-date:  2026-12-31 -->", "2026-12-31"}, // extra whitespace tolerated
	}
	for _, c := range cases {
		if got := parseSpecDate(c.body); got != c.want {
			t.Errorf("parseSpecDate(%q) = %q, want %q", c.body, got, c.want)
		}
	}
}

func TestUpdateCheckState(t *testing.T) {
	setup(t)
	recent := func() bool {
		return time.Since(time.Unix(readCheckState().CheckedAt, 0)) < updateCheckInterval
	}
	if recent() {
		t.Error("a fresh environment should not look recently-checked")
	}
	writeCheckState(checkState{CheckedAt: time.Now().Unix(), NudgedVersion: "1.2.0"})
	if !recent() {
		t.Error("after writeCheckState the check should look recent")
	}
	if got := readCheckState().NudgedVersion; got != "1.2.0" {
		t.Errorf("NudgedVersion round-trip = %q, want 1.2.0", got)
	}
}
