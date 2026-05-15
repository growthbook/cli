// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/growthbook-cli/internal/mocktest"
)

func TestSDKConnectionsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"sdk-connections", "create",
			"--environment", "environment",
			"--language", "language",
			"--name", "name",
			"--allowed-custom-fields-in-metadata", "string",
			"--encrypt-payload=true",
			"--hash-secure-attributes=true",
			"--include-custom-fields-in-metadata=true",
			"--include-draft-experiments=true",
			"--include-experiment-names=true",
			"--include-project-id-in-metadata=true",
			"--include-redirect-experiments=true",
			"--include-rule-ids=true",
			"--include-tags-in-metadata=true",
			"--include-visual-experiments=true",
			"--project", "string",
			"--proxy-enabled=true",
			"--proxy-host", "proxyHost",
			"--remote-eval-enabled=true",
			"--saved-group-references-enabled=true",
			"--sdk-version", "sdkVersion",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"environment: environment\n" +
			"language: language\n" +
			"name: name\n" +
			"allowedCustomFieldsInMetadata:\n" +
			"  - string\n" +
			"encryptPayload: true\n" +
			"hashSecureAttributes: true\n" +
			"includeCustomFieldsInMetadata: true\n" +
			"includeDraftExperiments: true\n" +
			"includeExperimentNames: true\n" +
			"includeProjectIdInMetadata: true\n" +
			"includeRedirectExperiments: true\n" +
			"includeRuleIds: true\n" +
			"includeTagsInMetadata: true\n" +
			"includeVisualExperiments: true\n" +
			"projects:\n" +
			"  - string\n" +
			"proxyEnabled: true\n" +
			"proxyHost: proxyHost\n" +
			"remoteEvalEnabled: true\n" +
			"savedGroupReferencesEnabled: true\n" +
			"sdkVersion: sdkVersion\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--domain", "string",
			"sdk-connections", "create",
		)
	})
}

func TestSDKConnectionsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"sdk-connections", "retrieve",
			"--id", "id",
		)
	})
}

func TestSDKConnectionsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"sdk-connections", "update",
			"--id", "id",
			"--allowed-custom-fields-in-metadata", "string",
			"--encrypt-payload=true",
			"--environment", "environment",
			"--hash-secure-attributes=true",
			"--include-custom-fields-in-metadata=true",
			"--include-draft-experiments=true",
			"--include-experiment-names=true",
			"--include-project-id-in-metadata=true",
			"--include-redirect-experiments=true",
			"--include-rule-ids=true",
			"--include-tags-in-metadata=true",
			"--include-visual-experiments=true",
			"--language", "language",
			"--name", "name",
			"--project", "string",
			"--proxy-enabled=true",
			"--proxy-host", "proxyHost",
			"--remote-eval-enabled=true",
			"--saved-group-references-enabled=true",
			"--sdk-version", "sdkVersion",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"allowedCustomFieldsInMetadata:\n" +
			"  - string\n" +
			"encryptPayload: true\n" +
			"environment: environment\n" +
			"hashSecureAttributes: true\n" +
			"includeCustomFieldsInMetadata: true\n" +
			"includeDraftExperiments: true\n" +
			"includeExperimentNames: true\n" +
			"includeProjectIdInMetadata: true\n" +
			"includeRedirectExperiments: true\n" +
			"includeRuleIds: true\n" +
			"includeTagsInMetadata: true\n" +
			"includeVisualExperiments: true\n" +
			"language: language\n" +
			"name: name\n" +
			"projects:\n" +
			"  - string\n" +
			"proxyEnabled: true\n" +
			"proxyHost: proxyHost\n" +
			"remoteEvalEnabled: true\n" +
			"savedGroupReferencesEnabled: true\n" +
			"sdkVersion: sdkVersion\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--domain", "string",
			"sdk-connections", "update",
			"--id", "id",
		)
	})
}

func TestSDKConnectionsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"sdk-connections", "list",
			"--limit", "1",
			"--multi-org", "multiOrg",
			"--offset", "0",
			"--project-id", "projectId",
			"--with-proxy", "withProxy",
		)
	})
}

func TestSDKConnectionsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"sdk-connections", "delete",
			"--id", "id",
		)
	})
}

func TestSDKConnectionsLookupByKey(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"sdk-connections", "lookup-by-key",
			"--key", "key",
		)
	})
}
