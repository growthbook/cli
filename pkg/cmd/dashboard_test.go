// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/growthbook-cli/internal/mocktest"
)

func TestDashboardsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"dashboards", "create",
			"--block", "{content: content, description: description, title: title, type: markdown, snapshotId: snapshotId}",
			"--edit-level", "published",
			"--enable-auto-updates=true",
			"--share-level", "published",
			"--title", "title",
			"--experiment-id", "experimentId",
			"--project", "string",
			"--update-schedule", "{hours: 0, type: stale}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"blocks:\n" +
			"  - content: content\n" +
			"    description: description\n" +
			"    title: title\n" +
			"    type: markdown\n" +
			"    snapshotId: snapshotId\n" +
			"editLevel: published\n" +
			"enableAutoUpdates: true\n" +
			"shareLevel: published\n" +
			"title: title\n" +
			"experimentId: experimentId\n" +
			"projects:\n" +
			"  - string\n" +
			"updateSchedule:\n" +
			"  hours: 0\n" +
			"  type: stale\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--domain", "string",
			"dashboards", "create",
		)
	})
}

func TestDashboardsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"dashboards", "retrieve",
			"--id", "id",
		)
	})
}

func TestDashboardsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"dashboards", "update",
			"--id", "id",
			"--block", "{content: content, description: description, title: title, type: markdown, snapshotId: snapshotId}",
			"--edit-level", "published",
			"--enable-auto-updates=true",
			"--project", "string",
			"--share-level", "published",
			"--title", "title",
			"--update-schedule", "{hours: 0, type: stale}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"blocks:\n" +
			"  - content: content\n" +
			"    description: description\n" +
			"    title: title\n" +
			"    type: markdown\n" +
			"    snapshotId: snapshotId\n" +
			"editLevel: published\n" +
			"enableAutoUpdates: true\n" +
			"projects:\n" +
			"  - string\n" +
			"shareLevel: published\n" +
			"title: title\n" +
			"updateSchedule:\n" +
			"  hours: 0\n" +
			"  type: stale\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--domain", "string",
			"dashboards", "update",
			"--id", "id",
		)
	})
}

func TestDashboardsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"dashboards", "list",
		)
	})
}

func TestDashboardsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"dashboards", "delete",
			"--id", "id",
		)
	})
}

func TestDashboardsListByExperiment(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"dashboards", "list-by-experiment",
			"--experiment-id", "experimentId",
		)
	})
}
