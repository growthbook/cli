// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/growthbook-cli/internal/mocktest"
	"github.com/stainless-sdks/growthbook-cli/internal/requestflag"
)

func TestFeaturesRevisionsRulesRampScheduleUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"features:revisions:rules:ramp-schedule", "update",
			"--id", "id",
			"--version", "new",
			"--rule-id", "ruleId",
			"--end-action", "{patch: {condition: condition, coverage: 0, enabled: true, force: {}, prerequisites: [{id: id, condition: condition}], ruleId: ruleId, savedGroups: [{ids: [string], match: all}]}, targetId: targetId, targetType: feature-rule}",
			"--end-condition", "{trigger: {at: {}, type: scheduled}}",
			"--environment", "environment",
			"--name", "name",
			"--revision-comment", "revisionComment",
			"--revision-title", "revisionTitle",
			"--start-date", "startDate",
			"--step", "{trigger: {seconds: 1, type: interval}, actions: [{patch: {condition: condition, coverage: 0, enabled: true, force: {}, prerequisites: [{id: id, condition: condition}], ruleId: ruleId, savedGroups: [{ids: [string], match: all}]}, targetId: targetId, targetType: feature-rule}], approvalNotes: approvalNotes}",
			"--template-id", "templateId",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(featuresRevisionsRulesRampScheduleUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"features:revisions:rules:ramp-schedule", "update",
			"--id", "id",
			"--version", "new",
			"--rule-id", "ruleId",
			"--end-action.patch", "{condition: condition, coverage: 0, enabled: true, force: {}, prerequisites: [{id: id, condition: condition}], ruleId: ruleId, savedGroups: [{ids: [string], match: all}]}",
			"--end-action.target-id", "targetId",
			"--end-action.target-type", "feature-rule",
			"--end-condition.trigger", "{at: {}, type: scheduled}",
			"--environment", "environment",
			"--name", "name",
			"--revision-comment", "revisionComment",
			"--revision-title", "revisionTitle",
			"--start-date", "startDate",
			"--step.trigger", "{seconds: 1, type: interval}",
			"--step.actions", "[{patch: {condition: condition, coverage: 0, enabled: true, force: {}, prerequisites: [{id: id, condition: condition}], ruleId: ruleId, savedGroups: [{ids: [string], match: all}]}, targetId: targetId, targetType: feature-rule}]",
			"--step.approval-notes", "approvalNotes",
			"--template-id", "templateId",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
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
			"    at: {}\n" +
			"    type: scheduled\n" +
			"environment: environment\n" +
			"name: name\n" +
			"revisionComment: revisionComment\n" +
			"revisionTitle: revisionTitle\n" +
			"startDate: startDate\n" +
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
			"features:revisions:rules:ramp-schedule", "update",
			"--id", "id",
			"--version", "new",
			"--rule-id", "ruleId",
		)
	})
}

func TestFeaturesRevisionsRulesRampScheduleDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"features:revisions:rules:ramp-schedule", "delete",
			"--id", "id",
			"--version", "new",
			"--rule-id", "ruleId",
			"--revision-comment", "revisionComment",
			"--revision-title", "revisionTitle",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"revisionComment: revisionComment\n" +
			"revisionTitle: revisionTitle\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"features:revisions:rules:ramp-schedule", "delete",
			"--id", "id",
			"--version", "new",
			"--rule-id", "ruleId",
		)
	})
}
