package customcfg

import "testing"

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
}

func TestUpdateCheckStamp(t *testing.T) {
	setup(t)
	if recentlyChecked() {
		t.Error("a fresh environment should not be recentlyChecked")
	}
	touchCheckStamp()
	if !recentlyChecked() {
		t.Error("after touchCheckStamp, recentlyChecked should be true")
	}
}
