// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/growthbook/cli/internal/mocktest"
	"github.com/growthbook/cli/internal/requestflag"
)

func TestRampScheduleTemplatesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ramp-schedule-templates", "create",
			"--name", "name",
			"--step", "{actions: [{patch: {ruleId: ruleId, condition: condition, coverage: 0, enabled: true, prerequisites: [{id: id, condition: condition}], savedGroups: [{ids: [string], match: all}]}, targetId: targetId, targetType: feature-rule}], trigger: {seconds: 1, type: interval}, approvalNotes: approvalNotes}",
			"--end-patch", "{condition: condition, coverage: 0, prerequisites: [{id: id, condition: condition}], savedGroups: [{ids: [string], match: all}]}",
			"--official=true",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(rampScheduleTemplatesCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ramp-schedule-templates", "create",
			"--name", "name",
			"--step.actions", "[{patch: {ruleId: ruleId, condition: condition, coverage: 0, enabled: true, prerequisites: [{id: id, condition: condition}], savedGroups: [{ids: [string], match: all}]}, targetId: targetId, targetType: feature-rule}]",
			"--step.trigger", "{seconds: 1, type: interval}",
			"--step.approval-notes", "approvalNotes",
			"--end-patch.condition", "condition",
			"--end-patch.coverage", "0",
			"--end-patch.prerequisites", "[{id: id, condition: condition}]",
			"--end-patch.saved-groups", "[{ids: [string], match: all}]",
			"--official=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: name\n" +
			"steps:\n" +
			"  - actions:\n" +
			"      - patch:\n" +
			"          ruleId: ruleId\n" +
			"          condition: condition\n" +
			"          coverage: 0\n" +
			"          enabled: true\n" +
			"          prerequisites:\n" +
			"            - id: id\n" +
			"              condition: condition\n" +
			"          savedGroups:\n" +
			"            - ids:\n" +
			"                - string\n" +
			"              match: all\n" +
			"        targetId: targetId\n" +
			"        targetType: feature-rule\n" +
			"    trigger:\n" +
			"      seconds: 1\n" +
			"      type: interval\n" +
			"    approvalNotes: approvalNotes\n" +
			"endPatch:\n" +
			"  condition: condition\n" +
			"  coverage: 0\n" +
			"  prerequisites:\n" +
			"    - id: id\n" +
			"      condition: condition\n" +
			"  savedGroups:\n" +
			"    - ids:\n" +
			"        - string\n" +
			"      match: all\n" +
			"official: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ramp-schedule-templates", "create",
		)
	})
}

func TestRampScheduleTemplatesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ramp-schedule-templates", "retrieve",
			"--id", "id",
		)
	})
}

func TestRampScheduleTemplatesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ramp-schedule-templates", "update",
			"--id", "id",
			"--end-patch", "{condition: condition, coverage: 0, prerequisites: [{id: id, condition: condition}], savedGroups: [{ids: [string], match: all}]}",
			"--name", "name",
			"--official=true",
			"--step", "{actions: [{patch: {ruleId: ruleId, condition: condition, coverage: 0, enabled: true, prerequisites: [{id: id, condition: condition}], savedGroups: [{ids: [string], match: all}]}, targetId: targetId, targetType: feature-rule}], trigger: {seconds: 1, type: interval}, approvalNotes: approvalNotes}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(rampScheduleTemplatesUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ramp-schedule-templates", "update",
			"--id", "id",
			"--end-patch.condition", "condition",
			"--end-patch.coverage", "0",
			"--end-patch.prerequisites", "[{id: id, condition: condition}]",
			"--end-patch.saved-groups", "[{ids: [string], match: all}]",
			"--name", "name",
			"--official=true",
			"--step.actions", "[{patch: {ruleId: ruleId, condition: condition, coverage: 0, enabled: true, prerequisites: [{id: id, condition: condition}], savedGroups: [{ids: [string], match: all}]}, targetId: targetId, targetType: feature-rule}]",
			"--step.trigger", "{seconds: 1, type: interval}",
			"--step.approval-notes", "approvalNotes",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"endPatch:\n" +
			"  condition: condition\n" +
			"  coverage: 0\n" +
			"  prerequisites:\n" +
			"    - id: id\n" +
			"      condition: condition\n" +
			"  savedGroups:\n" +
			"    - ids:\n" +
			"        - string\n" +
			"      match: all\n" +
			"name: name\n" +
			"official: true\n" +
			"steps:\n" +
			"  - actions:\n" +
			"      - patch:\n" +
			"          ruleId: ruleId\n" +
			"          condition: condition\n" +
			"          coverage: 0\n" +
			"          enabled: true\n" +
			"          prerequisites:\n" +
			"            - id: id\n" +
			"              condition: condition\n" +
			"          savedGroups:\n" +
			"            - ids:\n" +
			"                - string\n" +
			"              match: all\n" +
			"        targetId: targetId\n" +
			"        targetType: feature-rule\n" +
			"    trigger:\n" +
			"      seconds: 1\n" +
			"      type: interval\n" +
			"    approvalNotes: approvalNotes\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ramp-schedule-templates", "update",
			"--id", "id",
		)
	})
}

func TestRampScheduleTemplatesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ramp-schedule-templates", "list",
		)
	})
}

func TestRampScheduleTemplatesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ramp-schedule-templates", "delete",
			"--id", "id",
		)
	})
}
