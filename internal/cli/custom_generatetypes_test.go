package cli

import (
	"net/http"
	"os"
	"path/filepath"
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

func TestRenderAppAttributesTS(t *testing.T) {
	enum := "free, pro,enterprise"
	hash := true
	attributes := []components.Attribute{
		{Property: "plan", Datatype: components.AttributeDatatypeEnum, Enum: &enum},
		{Property: "id", Datatype: components.AttributeDatatypeString},
		{Property: "company.id", Datatype: components.AttributeDatatypeString, HashAttribute: &hash},
		{Property: "company", Datatype: components.AttributeDatatypeString},
		{Property: "company.size", Datatype: components.AttributeDatatypeNumber},
		{Property: "company.name", Datatype: components.AttributeDatatypeSecureString},
		{Property: "tags", Datatype: components.AttributeDatatypeStringArray, Enum: &enum},
		{Property: "scores", Datatype: components.AttributeDatatypeNumberArray},
		{Property: "it's", Datatype: components.AttributeDatatypeBoolean},
		{Property: "status", Datatype: components.AttributeDatatypeEnum},
	}

	want := `export type AppAttributes = {
  'company'?: string | {
    'id'?: string | null;
    'name'?: string | null;
    'size'?: number | null;
  } | null;
  'company.id'?: string | null;
  'id'?: string | null;
  'it\'s'?: boolean | null;
  'plan'?: 'free' | 'pro' | 'enterprise' | null;
  'scores'?: number[] | null;
  'status'?: string | null;
  'tags'?: ('free' | 'pro' | 'enterprise')[] | null;
}
`
	if out := renderAppAttributesTS(attributes, false); out != want {
		t.Errorf("renderAppAttributesTS =\n%s\nwant:\n%s", out, want)
	}

	wantStrict := `export type AppAttributes = {
  'company': string | {
    'id': string | null | undefined;
    'name': string | null | undefined;
    'size': number | null | undefined;
  } | null | undefined;
  'company.id': string | null | undefined;
  'id': string | null | undefined;
  'it\'s': boolean | null | undefined;
  'plan': 'free' | 'pro' | 'enterprise' | null | undefined;
  'scores': number[] | null | undefined;
  'status': string | null | undefined;
  'tags': ('free' | 'pro' | 'enterprise')[] | null | undefined;
}
`
	if out := renderAppAttributesTS(attributes, true); out != wantStrict {
		t.Errorf("strict renderAppAttributesTS =\n%s\nwant:\n%s", out, wantStrict)
	}
}

func TestE2E_GenerateTypes(t *testing.T) {
	var featureProjects []string
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case strings.HasSuffix(r.URL.Path, "/attributes"):
			_, _ = w.Write([]byte(`{"attributes":[
				{"property":"id","datatype":"string","hashAttribute":true},
				{"property":"old","datatype":"string","archived":true},
				{"property":"plan","datatype":"string","projects":["prj_2","prj_9"]},
				{"property":"other","datatype":"string","projects":["prj_9"]}
			]}`))
		case strings.HasSuffix(r.URL.Path, "/features"):
			project := r.URL.Query().Get("projectId")
			featureProjects = append(featureProjects, project)
			if project == "prj_empty" {
				_, _ = w.Write([]byte(`{"features":[],"limit":100,"offset":0,"count":0,"total":0,"hasMore":false,"nextOffset":null}`))
				return
			}
			_, _ = w.Write([]byte(`{"features":[{"id":"flag-` + project + `","valueType":"boolean","defaultValue":"false","project":"` + project + `"}],"limit":100,"offset":0,"count":1,"total":1,"hasMore":false,"nextOffset":null}`))
		default:
			t.Errorf("unexpected request: %s", r.URL.Path)
		}
	}

	// Several projects, given both ways the flag allows, with one repeated and
	// a trailing comma that must not add an unfiltered request.
	dir := t.TempDir()
	out, stderr, err := runCLI(t, handler, "generate-types", "--output", dir, "--project", "prj_1,prj_2,", "--project", "prj_1")
	if err != nil {
		t.Fatalf("generate-types returned error: %v", err)
	}
	if got := strings.Join(featureProjects, ","); got != "prj_1,prj_2,prj_1" {
		t.Errorf("features requested for projects %q, want one request per --project", got)
	}
	if strings.Contains(stderr, "Warning") {
		t.Errorf("unexpected warning: %s", stderr)
	}
	if !strings.Contains(out, "2 feature(s) and 2 attribute(s)") {
		t.Errorf("unexpected stdout: %s", out)
	}
	got, err := os.ReadFile(filepath.Join(dir, defaultTypesFilename))
	if err != nil {
		t.Fatal(err)
	}
	want := `export type AppFeatures = {
  'flag-prj_1': boolean;
  'flag-prj_2': boolean;
}

export type AppAttributes = {
  'id'?: string | null;
  'plan'?: string | null;
}

// @ts-ignore - only applies if your project imports this package
declare module '@growthbook/growthbook' {
  interface Register {
    features: AppFeatures;
    attributes: AppAttributes;
  }
}

// @ts-ignore - only applies if your project imports this package
declare module '@growthbook/growthbook-react' {
  interface Register {
    features: AppFeatures;
    attributes: AppAttributes;
  }
}
`
	if !strings.HasSuffix(string(got), want) {
		t.Errorf("generated file does not end with:\n%s\ngot:\n%s", want, got)
	}

	// No --project: one unfiltered request and every unarchived attribute.
	// --strict-attributes requires each key, and --no-register leaves the
	// augmentation out.
	featureProjects = nil
	out, _, err = runCLI(t, handler, "generate-types", "--output", dir, "--no-register", "--strict-attributes")
	if err != nil {
		t.Fatalf("generate-types --no-register returned error: %v", err)
	}
	if got := strings.Join(featureProjects, ","); got != "" || len(featureProjects) != 1 {
		t.Errorf("features requested for projects %q, want a single unfiltered request", featureProjects)
	}
	if !strings.Contains(out, "1 feature(s) and 3 attribute(s)") {
		t.Errorf("unexpected stdout: %s", out)
	}
	got, _ = os.ReadFile(filepath.Join(dir, defaultTypesFilename))
	if !strings.Contains(string(got), "  'id': string | null | undefined;\n") {
		t.Errorf("--strict-attributes did not require the keys:\n%s", got)
	}
	if strings.Contains(string(got), "declare module") {
		t.Errorf("--no-register still registered the types:\n%s", got)
	}

	// Registering zero features makes every feature key a type error, so warn.
	_, stderr, err = runCLI(t, handler, "generate-types", "--output", dir, "--project", "prj_empty")
	if err != nil {
		t.Fatalf("generate-types --project prj_empty returned error: %v", err)
	}
	if !strings.Contains(stderr, "Warning: no features found") {
		t.Errorf("expected a no-features warning, got stderr: %s", stderr)
	}
}
