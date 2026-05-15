// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/growthbook-cli/internal/mocktest"
	"github.com/stainless-sdks/growthbook-cli/internal/requestflag"
)

func TestRampSchedulesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ramp-schedules", "create",
			"--name", "name",
			"--end-action", "{patch: {condition: condition, coverage: 0, enabled: true, force: {}, prerequisites: [{id: id, condition: condition}], ruleId: ruleId, savedGroups: [{ids: [string], match: all}]}, targetId: targetId, targetType: feature-rule}",
			"--end-condition", "{trigger: {at: '2019-12-27T18:11:19.117Z', type: scheduled}}",
			"--environment", "environment",
			"--feature-id", "featureId",
			"--rule-id", "ruleId",
			"--start-date", "'2019-12-27T18:11:19.117Z'",
			"--step", "{trigger: {seconds: 1, type: interval}, actions: [{patch: {condition: condition, coverage: 0, enabled: true, force: {}, prerequisites: [{id: id, condition: condition}], ruleId: ruleId, savedGroups: [{ids: [string], match: all}]}, targetId: targetId, targetType: feature-rule}], approvalNotes: approvalNotes}",
			"--template-id", "templateId",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(rampSchedulesCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ramp-schedules", "create",
			"--name", "name",
			"--end-action.patch", "{condition: condition, coverage: 0, enabled: true, force: {}, prerequisites: [{id: id, condition: condition}], ruleId: ruleId, savedGroups: [{ids: [string], match: all}]}",
			"--end-action.target-id", "targetId",
			"--end-action.target-type", "feature-rule",
			"--end-condition.trigger", "{at: '2019-12-27T18:11:19.117Z', type: scheduled}",
			"--environment", "environment",
			"--feature-id", "featureId",
			"--rule-id", "ruleId",
			"--start-date", "'2019-12-27T18:11:19.117Z'",
			"--step.trigger", "{seconds: 1, type: interval}",
			"--step.actions", "[{patch: {condition: condition, coverage: 0, enabled: true, force: {}, prerequisites: [{id: id, condition: condition}], ruleId: ruleId, savedGroups: [{ids: [string], match: all}]}, targetId: targetId, targetType: feature-rule}]",
			"--step.approval-notes", "approvalNotes",
			"--template-id", "templateId",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: name\n" +
			"endActions:\n" +
			"  - patch:\n" +
			"      condition: condition\n" +
			"      coverage: 0\n" +
			"      enabled: true\n" +
			"      force: {}\n" +
			"      prerequisites:\n" +
			"        - id: id\n" +
			"          condition: condition\n" +
			"      ruleId: ruleId\n" +
			"      savedGroups:\n" +
			"        - ids:\n" +
			"            - string\n" +
			"          match: all\n" +
			"    targetId: targetId\n" +
			"    targetType: feature-rule\n" +
			"endCondition:\n" +
			"  trigger:\n" +
			"    at: '2019-12-27T18:11:19.117Z'\n" +
			"    type: scheduled\n" +
			"environment: environment\n" +
			"featureId: featureId\n" +
			"ruleId: ruleId\n" +
			"startDate: '2019-12-27T18:11:19.117Z'\n" +
			"steps:\n" +
			"  - trigger:\n" +
			"      seconds: 1\n" +
			"      type: interval\n" +
			"    actions:\n" +
			"      - patch:\n" +
			"          condition: condition\n" +
			"          coverage: 0\n" +
			"          enabled: true\n" +
			"          force: {}\n" +
			"          prerequisites:\n" +
			"            - id: id\n" +
			"              condition: condition\n" +
			"          ruleId: ruleId\n" +
			"          savedGroups:\n" +
			"            - ids:\n" +
			"                - string\n" +
			"              match: all\n" +
			"        targetId: targetId\n" +
			"        targetType: feature-rule\n" +
			"    approvalNotes: approvalNotes\n" +
			"templateId: templateId\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ramp-schedules", "create",
		)
	})
}

func TestRampSchedulesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ramp-schedules", "retrieve",
			"--id", "id",
		)
	})
}

func TestRampSchedulesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ramp-schedules", "update",
			"--id", "id",
			"--end-action", "{patch: {ruleId: ruleId, condition: condition, coverage: 0, enabled: true, force: {}, prerequisites: [{id: id, condition: condition}], savedGroups: [{ids: [string], match: all}]}, targetId: targetId, targetType: feature-rule}",
			"--end-condition", "{trigger: {at: '2019-12-27T18:11:19.117Z', type: scheduled}}",
			"--name", "name",
			"--start-date", "'2019-12-27T18:11:19.117Z'",
			"--step", "{trigger: {seconds: 1, type: interval}, actions: [{patch: {ruleId: ruleId, condition: condition, coverage: 0, enabled: true, force: {}, prerequisites: [{id: id, condition: condition}], savedGroups: [{ids: [string], match: all}]}, targetId: targetId, targetType: feature-rule}], approvalNotes: approvalNotes}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(rampSchedulesUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ramp-schedules", "update",
			"--id", "id",
			"--end-action.patch", "{ruleId: ruleId, condition: condition, coverage: 0, enabled: true, force: {}, prerequisites: [{id: id, condition: condition}], savedGroups: [{ids: [string], match: all}]}",
			"--end-action.target-id", "targetId",
			"--end-action.target-type", "feature-rule",
			"--end-condition.trigger", "{at: '2019-12-27T18:11:19.117Z', type: scheduled}",
			"--name", "name",
			"--start-date", "'2019-12-27T18:11:19.117Z'",
			"--step.trigger", "{seconds: 1, type: interval}",
			"--step.actions", "[{patch: {ruleId: ruleId, condition: condition, coverage: 0, enabled: true, force: {}, prerequisites: [{id: id, condition: condition}], savedGroups: [{ids: [string], match: all}]}, targetId: targetId, targetType: feature-rule}]",
			"--step.approval-notes", "approvalNotes",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"endActions:\n" +
			"  - patch:\n" +
			"      ruleId: ruleId\n" +
			"      condition: condition\n" +
			"      coverage: 0\n" +
			"      enabled: true\n" +
			"      force: {}\n" +
			"      prerequisites:\n" +
			"        - id: id\n" +
			"          condition: condition\n" +
			"      savedGroups:\n" +
			"        - ids:\n" +
			"            - string\n" +
			"          match: all\n" +
			"    targetId: targetId\n" +
			"    targetType: feature-rule\n" +
			"endCondition:\n" +
			"  trigger:\n" +
			"    at: '2019-12-27T18:11:19.117Z'\n" +
			"    type: scheduled\n" +
			"name: name\n" +
			"startDate: '2019-12-27T18:11:19.117Z'\n" +
			"steps:\n" +
			"  - trigger:\n" +
			"      seconds: 1\n" +
			"      type: interval\n" +
			"    actions:\n" +
			"      - patch:\n" +
			"          ruleId: ruleId\n" +
			"          condition: condition\n" +
			"          coverage: 0\n" +
			"          enabled: true\n" +
			"          force: {}\n" +
			"          prerequisites:\n" +
			"            - id: id\n" +
			"              condition: condition\n" +
			"          savedGroups:\n" +
			"            - ids:\n" +
			"                - string\n" +
			"              match: all\n" +
			"        targetId: targetId\n" +
			"        targetType: feature-rule\n" +
			"    approvalNotes: approvalNotes\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ramp-schedules", "update",
			"--id", "id",
		)
	})
}

func TestRampSchedulesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ramp-schedules", "list",
			"--feature-id", "featureId",
			"--limit", "1",
			"--offset", "0",
			"--status", "pending",
		)
	})
}

func TestRampSchedulesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ramp-schedules", "delete",
			"--id", "id",
		)
	})
}
