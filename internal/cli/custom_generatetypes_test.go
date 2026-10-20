package cli

import (
	"strings"
	"testing"

	"github.com/growthbook/cli/v3/internal/sdk/models/components"
)

func TestTSType(t *testing.T) {
	cases := []struct {
		vt           components.FeatureV2ValueType
		defaultValue string
		want         string
	}{
		{components.FeatureV2ValueTypeBoolean, "true", "boolean"},
		{components.FeatureV2ValueTypeNumber, "1", "number"},
		{components.FeatureV2ValueTypeString, "[x]", "string"},
		{components.FeatureV2ValueTypeJSON, `{"a":1}`, "Record<string, unknown>"},
		{components.FeatureV2ValueTypeJSON, " [\"us\"]", "unknown[]"},
		{components.FeatureV2ValueTypeJSON, "", "Record<string, unknown>"},
		{"unexpected", "", "unknown"},
	}
	for _, c := range cases {
		f := components.FeatureV2{ValueType: c.vt, DefaultValue: c.defaultValue}
		if got := tsType(f); got != c.want {
			t.Errorf("tsType(%q, %q) = %q, want %q", c.vt, c.defaultValue, got, c.want)
		}
	}
}

func TestRenderAppFeaturesTSSortsAndTypesFeatures(t *testing.T) {
	out := renderAppFeaturesTS([]components.FeatureV2{
		{ID: "c-markets", ValueType: components.FeatureV2ValueTypeJSON, DefaultValue: "[]"},
		{ID: "b-config", ValueType: components.FeatureV2ValueTypeJSON, DefaultValue: "{}"},
		{ID: "a-enabled", ValueType: components.FeatureV2ValueTypeBoolean},
	})
	want := "export type AppFeatures = {\n" +
		"  'a-enabled': boolean;\n" +
		"  'b-config': Record<string, unknown>;\n" +
		"  'c-markets': unknown[];\n" +
		"}\n"
	if !strings.HasSuffix(out, want) {
		t.Errorf("renderAppFeaturesTS output does not end with:\n%s\ngot:\n%s", want, out)
	}
}
