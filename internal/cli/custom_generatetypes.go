// Hand-written (NOT generated). Ported from the legacy TS CLI's
// `features generate-types`; preserved across regen via persistentEdits.
package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strings"

	"github.com/growthbook/cli/v3/internal/client"
	"github.com/growthbook/cli/v3/internal/sdk"
	"github.com/growthbook/cli/v3/internal/sdk/models/components"
	"github.com/growthbook/cli/v3/internal/sdk/models/operations"
	"github.com/spf13/cobra"
)

const (
	defaultTypesDir      = "growthbook-types"
	defaultTypesFilename = "app-features.ts"
)

// sdkPackages are the SDK packages that export the `Register` interface.
var sdkPackages = []string{"@growthbook/growthbook", "@growthbook/growthbook-react"}

// newGenerateTypesCmd builds the `generate-types` command, which fetches all
// features and attributes and writes TypeScript `AppFeatures` and
// `AppAttributes` definitions for type-safe SDK usage. (Legacy CLI exposed this
// as `features generate-types`.)
func newGenerateTypesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate-types",
		Short: "Generate TypeScript types for all your features and attributes",
		Long: "Fetches all of your features and attributes and writes TypeScript `AppFeatures` and " +
			"`AppAttributes` type definitions, giving you type-safe feature keys, feature values and " +
			"targeting attributes in the GrowthBook SDK.\n\n" +
			"The file also registers the types with the SDK, so every GrowthBook instance and React " +
			"hook is typed with no generics (JS and React SDK 2.1.0+).",
		Example: "  growthbook generate-types\n" +
			"  growthbook generate-types --output ./src/types --filename growthbook.ts\n" +
			"  growthbook generate-types --project prj_123\n" +
			"  growthbook generate-types --project prj_123 --project prj_456\n" +
			"  growthbook generate-types --strict-attributes\n" +
			"  growthbook generate-types --no-register",
		RunE: runGenerateTypes,
	}
	cmd.Flags().String("output", "", fmt.Sprintf("Output directory for the types file (default ./%s)", defaultTypesDir))
	cmd.Flags().StringP("filename", "f", defaultTypesFilename, "Output filename for the generated types")
	cmd.Flags().StringSliceP("project", "p", nil, "Filter features and attributes by project id. Repeat the flag or comma-separate ids for several projects, e.g. to match an SDK Connection")
	cmd.Flags().Bool("strict-attributes", false, "Make every attribute a required key, so leaving one out when setting attributes is a type error")
	cmd.Flags().Bool("no-register", false, "Only write the type definitions, without registering them with the SDK")
	return cmd
}

func runGenerateTypes(cmd *cobra.Command, args []string) error {
	outputDir, _ := cmd.Flags().GetString("output")
	filename, _ := cmd.Flags().GetString("filename")
	projects, _ := cmd.Flags().GetStringSlice("project")
	// An empty id (e.g. a trailing comma) would otherwise fetch every feature unfiltered.
	projects = slices.DeleteFunc(projects, func(p string) bool { return p == "" })
	strictAttributes, _ := cmd.Flags().GetBool("strict-attributes")
	noRegister, _ := cmd.Flags().GetBool("no-register")

	s, err := client.NewClient(cmd)
	if err != nil {
		return err
	}

	features, err := fetchAllFeatures(cmd, s, projects)
	if err != nil {
		return err
	}
	attributes, err := fetchAttributes(cmd, s, projects)
	if err != nil {
		return err
	}

	if outputDir == "" {
		outputDir = defaultTypesDir
	}
	if !strings.HasSuffix(filename, ".ts") {
		filename += ".ts"
	}
	if err := os.MkdirAll(outputDir, 0o755); err != nil {
		return fmt.Errorf("create output directory %q: %w", outputDir, err)
	}
	content := renderAppFeaturesTS(features) + "\n" + renderAppAttributesTS(attributes, strictAttributes)
	if !noRegister {
		content += renderRegisterTS()
	}
	outPath := filepath.Join(outputDir, filename)
	if err := os.WriteFile(outPath, []byte(content), 0o644); err != nil {
		return fmt.Errorf("write types file %q: %w", outPath, err)
	}

	fmt.Fprintf(cmd.OutOrStdout(), "Wrote TypeScript definitions for %d feature(s) and %d attribute(s) to %s\n", len(features), len(attributes), outPath)
	if len(features) == 0 && !noRegister {
		fmt.Fprintln(cmd.ErrOrStderr(), "Warning: no features found, so every feature key in your app will be a type error. Check --project and your API key's access, or pass --no-register.")
	}
	return nil
}

// fetchAttributes lists the targeting attributes, minus archived ones. With
// projects, it keeps the org-wide attributes plus the ones scoped to any of
// them. That filter runs here because the API takes a single project, and
// GrowthBook servers before 5.0 ignore it.
func fetchAttributes(cmd *cobra.Command, s *sdk.Growthbook, projects []string) ([]components.Attribute, error) {
	res, err := s.Attributes.ListAttributes(cmd.Context(), &operations.ListAttributesRequest{})
	if err != nil {
		return nil, err
	}
	var attributes []components.Attribute
	for _, a := range res.GetObject().GetAttributes() {
		if archived := a.GetArchived(); archived != nil && *archived {
			continue
		}
		if len(projects) == 0 || len(a.GetProjects()) == 0 || slices.ContainsFunc(a.GetProjects(), func(p string) bool {
			return slices.Contains(projects, p)
		}) {
			attributes = append(attributes, a)
		}
	}
	return attributes, nil
}

// fetchAllFeatures pages through the v2 features endpoint until exhausted,
// once per project when filtering (the API takes a single project).
func fetchAllFeatures(cmd *cobra.Command, s *sdk.Growthbook, projects []string) ([]components.FeatureV2, error) {
	if len(projects) == 0 {
		projects = []string{""}
	}
	var (
		features []components.FeatureV2
		seen           = map[string]bool{}
		limit    int64 = 100
	)
	for _, project := range projects {
		var offset int64
		for {
			req := operations.ListFeaturesV2Request{Limit: &limit, Offset: &offset}
			if project != "" {
				req.ProjectID = &project
			}
			res, err := s.Features.List(cmd.Context(), &req)
			if err != nil {
				return nil, err
			}
			if res.Object == nil {
				break
			}
			// A project listed twice would otherwise repeat its keys in the type.
			for _, f := range res.Object.Features {
				if !seen[f.GetID()] {
					seen[f.GetID()] = true
					features = append(features, f)
				}
			}
			if !res.Object.HasMore || res.Object.NextOffset == nil {
				break
			}
			offset = *res.Object.NextOffset
		}
	}
	return features, nil
}

// renderAppFeaturesTS renders the AppFeatures type. Features are sorted by id
// so the output is deterministic (stable diffs when committed to a repo).
func renderAppFeaturesTS(features []components.FeatureV2) string {
	sort.Slice(features, func(i, j int) bool { return features[i].GetID() < features[j].GetID() })

	var b strings.Builder
	b.WriteString(`/* eslint-disable */
/**
 * GrowthBook TypeScript Typings
 *
 * NOTE: This file is auto generated by the GrowthBook CLI (https://docs.growthbook.io/tools/cli).
 * See https://docs.growthbook.io/tools/cli#generating-typescript-types for more information.
 *
 * DO NOT EDIT this file manually.
 */
export type AppFeatures = {
`)
	for _, f := range features {
		fmt.Fprintf(&b, "  '%s': %s;\n", f.GetID(), tsType(f))
	}
	b.WriteString("}\n")
	return b.String()
}

func tsType(f components.FeatureV2) string {
	switch vt := f.GetValueType(); vt {
	case components.FeatureV2ValueTypeBoolean, components.FeatureV2ValueTypeNumber, components.FeatureV2ValueTypeString:
		return string(vt)
	case components.FeatureV2ValueTypeJSON:
		// Rule values must match the default's shape, so an array default means an array feature.
		if strings.HasPrefix(strings.TrimSpace(f.GetDefaultValue()), "[") {
			return "unknown[]"
		}
		return "Record<string, unknown>"
	default:
		return "unknown"
	}
}

// attributeNode is one key of AppAttributes. A dotted property such as
// `company.plan` is a nested object, because that is how the SDK resolves it.
type attributeNode struct {
	ts       string
	children map[string]*attributeNode
}

// renderAppAttributesTS renders the AppAttributes type.
//
// Keys are optional by default: apps commonly set attributes in stages, plugins
// fill some in, and an org's attributes are shared by apps that each send only
// a subset. Unknown keys and wrong value types are still compile errors.
//
// With strict, every key is required, so forgetting an attribute is a compile
// error too. Its value may still be null or undefined for when the app doesn't
// know it (e.g. the id of a logged-out user).
func renderAppAttributesTS(attributes []components.Attribute, strict bool) string {
	root := &attributeNode{}
	for _, a := range attributes {
		node := root
		for _, part := range strings.Split(a.GetProperty(), ".") {
			if part == "" {
				continue
			}
			if node.children == nil {
				node.children = map[string]*attributeNode{}
			}
			if node.children[part] == nil {
				node.children[part] = &attributeNode{}
			}
			node = node.children[part]
		}
		node.ts = attributeTSType(a)
		// The SDK reads a hash attribute as a flat key, so a dotted one needs that form too.
		if hash := a.GetHashAttribute(); hash != nil && *hash && strings.Contains(a.GetProperty(), ".") {
			if root.children == nil {
				root.children = map[string]*attributeNode{}
			}
			root.children[a.GetProperty()] = &attributeNode{ts: node.ts}
		}
	}

	var b strings.Builder
	b.WriteString("export type AppAttributes = ")
	writeAttributeNode(&b, root, "", strict)
	b.WriteString("\n")
	return b.String()
}

func writeAttributeNode(b *strings.Builder, node *attributeNode, indent string, strict bool) {
	// An optional key already accepts undefined.
	key, nullable := "%s  %s?: ", " | null;\n"
	if strict {
		key, nullable = "%s  %s: ", " | null | undefined;\n"
	}

	keys := make([]string, 0, len(node.children))
	for k := range node.children {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	b.WriteString("{\n")
	for _, k := range keys {
		child := node.children[k]
		fmt.Fprintf(b, key, indent, tsString(k))
		if child.children != nil {
			writeAttributeNode(b, child, indent+"  ", strict)
		} else {
			b.WriteString(child.ts)
		}
		b.WriteString(nullable)
	}
	b.WriteString(indent + "}")
}

func attributeTSType(a components.Attribute) string {
	// The allowed values of an enum, as a union of string literals.
	var values []string
	if enum := a.GetEnum(); enum != nil {
		for _, v := range strings.Split(*enum, ",") {
			if v = strings.TrimSpace(v); v != "" {
				values = append(values, tsString(v))
			}
		}
	}
	union := strings.Join(values, " | ")

	switch a.GetDatatype() {
	case components.AttributeDatatypeBoolean:
		return "boolean"
	case components.AttributeDatatypeNumber:
		return "number"
	case components.AttributeDatatypeNumberArray:
		return "number[]"
	case components.AttributeDatatypeString, components.AttributeDatatypeSecureString:
		return "string"
	case components.AttributeDatatypeEnum:
		if union != "" {
			return union
		}
		return "string"
	case components.AttributeDatatypeStringArray, components.AttributeDatatypeSecureStringArray:
		if union != "" {
			return "(" + union + ")[]"
		}
		return "string[]"
	default:
		return "unknown"
	}
}

// renderRegisterTS registers the types on the SDK's `Register` interface, which
// makes them the default for every GrowthBook instance and React hook.
//
// It registers with both SDK packages because which one a project can augment
// depends on what it installs and imports: augmenting a package the project
// doesn't import is a compile error (TS2664), which the @ts-ignore suppresses.
func renderRegisterTS() string {
	var b strings.Builder
	for _, pkg := range sdkPackages {
		fmt.Fprintf(&b, `
// @ts-ignore - only applies if your project imports this package
declare module %s {
  interface Register {
    features: AppFeatures;
    attributes: AppAttributes;
  }
}
`, tsString(pkg))
	}
	return b.String()
}

// tsString renders s as a single-quoted TypeScript string literal.
func tsString(s string) string {
	return "'" + strings.NewReplacer(`\`, `\\`, "'", `\'`, "\n", `\n`).Replace(s) + "'"
}
