// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/growthbook-cli/internal/mocktest"
	"github.com/stainless-sdks/growthbook-cli/internal/requestflag"
)

func TestFeaturesRevisionsRulesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"features:revisions:rules", "update",
			"--id", "id",
			"--version", "new",
			"--rule-id", "ruleId",
			"--rule", "{allEnvironments: true, condition: condition, controlValue: controlValue, coverage: 0, description: description, enabled: true, environments: [string], experimentId: experimentId, hashAttribute: hashAttribute, prerequisites: [{id: id, condition: condition}], savedGroups: [{ids: [string], match: all}], scheduleRules: [{enabled: true, timestamp: timestamp}], scheduleType: none, seed: seed, type: force, value: value, variations: [{value: value, variationId: variationId}], variationValue: variationValue}",
			"--ramp-schedule", "{endActions: [{patch: {condition: condition, coverage: 0, enabled: true, force: {}, prerequisites: [{id: id, condition: condition}], ruleId: ruleId, savedGroups: [{ids: [string], match: all}]}, targetId: targetId, targetType: feature-rule}], endCondition: {trigger: {at: {}, type: scheduled}}, name: name, startDate: startDate, steps: [{trigger: {seconds: 1, type: interval}, actions: [{patch: {condition: condition, coverage: 0, enabled: true, force: {}, prerequisites: [{id: id, condition: condition}], ruleId: ruleId, savedGroups: [{ids: [string], match: all}]}, targetId: targetId, targetType: feature-rule}], approvalNotes: approvalNotes}], templateId: templateId}",
			"--revision-comment", "revisionComment",
			"--revision-title", "revisionTitle",
			"--schedule", "{endDate: endDate, startDate: startDate}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(featuresRevisionsRulesUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"features:revisions:rules", "update",
			"--id", "id",
			"--version", "new",
			"--rule-id", "ruleId",
			"--rule.all-environments=true",
			"--rule.condition", "condition",
			"--rule.control-value", "controlValue",
			"--rule.coverage", "0",
			"--rule.description", "description",
			"--rule.enabled=true",
			"--rule.environments", "[string]",
			"--rule.experiment-id", "experimentId",
			"--rule.hash-attribute", "hashAttribute",
			"--rule.prerequisites", "[{id: id, condition: condition}]",
			"--rule.saved-groups", "[{ids: [string], match: all}]",
			"--rule.schedule-rules", "[{enabled: true, timestamp: timestamp}]",
			"--rule.schedule-type", "none",
			"--rule.seed", "seed",
			"--rule.type", "force",
			"--rule.value", "value",
			"--rule.variations", "[{value: value, variationId: variationId}]",
			"--rule.variation-value", "variationValue",
			"--ramp-schedule.end-actions", "[{patch: {condition: condition, coverage: 0, enabled: true, force: {}, prerequisites: [{id: id, condition: condition}], ruleId: ruleId, savedGroups: [{ids: [string], match: all}]}, targetId: targetId, targetType: feature-rule}]",
			"--ramp-schedule.end-condition", "{trigger: {at: {}, type: scheduled}}",
			"--ramp-schedule.name", "name",
			"--ramp-schedule.start-date", "startDate",
			"--ramp-schedule.steps", "[{trigger: {seconds: 1, type: interval}, actions: [{patch: {condition: condition, coverage: 0, enabled: true, force: {}, prerequisites: [{id: id, condition: condition}], ruleId: ruleId, savedGroups: [{ids: [string], match: all}]}, targetId: targetId, targetType: feature-rule}], approvalNotes: approvalNotes}]",
			"--ramp-schedule.template-id", "templateId",
			"--revision-comment", "revisionComment",
			"--revision-title", "revisionTitle",
			"--schedule.end-date", "endDate",
			"--schedule.start-date", "startDate",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"rule:\n" +
			"  allEnvironments: true\n" +
			"  condition: condition\n" +
			"  controlValue: controlValue\n" +
			"  coverage: 0\n" +
			"  description: description\n" +
			"  enabled: true\n" +
			"  environments:\n" +
			"    - string\n" +
			"  experimentId: experimentId\n" +
			"  hashAttribute: hashAttribute\n" +
			"  prerequisites:\n" +
			"    - id: id\n" +
			"      condition: condition\n" +
			"  savedGroups:\n" +
			"    - ids:\n" +
			"        - string\n" +
			"      match: all\n" +
			"  scheduleRules:\n" +
			"    - enabled: true\n" +
			"      timestamp: timestamp\n" +
			"  scheduleType: none\n" +
			"  seed: seed\n" +
			"  type: force\n" +
			"  value: value\n" +
			"  variations:\n" +
			"    - value: value\n" +
			"      variationId: variationId\n" +
			"  variationValue: variationValue\n" +
			"rampSchedule:\n" +
			"  endActions:\n" +
			"    - patch:\n" +
			"        condition: condition\n" +
			"        coverage: 0\n" +
			"        enabled: true\n" +
			"        force: {}\n" +
			"        prerequisites:\n" +
			"          - id: id\n" +
			"            condition: condition\n" +
			"        ruleId: ruleId\n" +
			"        savedGroups:\n" +
			"          - ids:\n" +
			"              - string\n" +
			"            match: all\n" +
			"      targetId: targetId\n" +
			"      targetType: feature-rule\n" +
			"  endCondition:\n" +
			"    trigger:\n" +
			"      at: {}\n" +
			"      type: scheduled\n" +
			"  name: name\n" +
			"  startDate: startDate\n" +
			"  steps:\n" +
			"    - trigger:\n" +
			"        seconds: 1\n" +
			"        type: interval\n" +
			"      actions:\n" +
			"        - patch:\n" +
			"            condition: condition\n" +
			"            coverage: 0\n" +
			"            enabled: true\n" +
			"            force: {}\n" +
			"            prerequisites:\n" +
			"              - id: id\n" +
			"                condition: condition\n" +
			"            ruleId: ruleId\n" +
			"            savedGroups:\n" +
			"              - ids:\n" +
			"                  - string\n" +
			"                match: all\n" +
			"          targetId: targetId\n" +
			"          targetType: feature-rule\n" +
			"      approvalNotes: approvalNotes\n" +
			"  templateId: templateId\n" +
			"revisionComment: revisionComment\n" +
			"revisionTitle: revisionTitle\n" +
			"schedule:\n" +
			"  endDate: endDate\n" +
			"  startDate: startDate\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"features:revisions:rules", "update",
			"--id", "id",
			"--version", "new",
			"--rule-id", "ruleId",
		)
	})
}

func TestFeaturesRevisionsRulesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"features:revisions:rules", "delete",
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
			"features:revisions:rules", "delete",
			"--id", "id",
			"--version", "new",
			"--rule-id", "ruleId",
		)
	})
}

func TestFeaturesRevisionsRulesAdd(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"features:revisions:rules", "add",
			"--id", "id",
			"--version", "new",
			"--rule", "{experimentId: experimentId, type: experiment-ref, variations: [{value: value, variationId: variationId}], allEnvironments: true, condition: condition, description: description, enabled: true, environments: [string], prerequisites: [{id: id, condition: condition}], savedGroups: [{ids: [string], match: all}], scheduleRules: [{enabled: true, timestamp: timestamp}], scheduleType: none}",
			"--ramp-schedule", "{endActions: [{patch: {condition: condition, coverage: 0, enabled: true, force: {}, prerequisites: [{id: id, condition: condition}], ruleId: ruleId, savedGroups: [{ids: [string], match: all}]}, targetId: targetId, targetType: feature-rule}], endCondition: {trigger: {at: {}, type: scheduled}}, name: name, startDate: startDate, steps: [{trigger: {seconds: 1, type: interval}, actions: [{patch: {condition: condition, coverage: 0, enabled: true, force: {}, prerequisites: [{id: id, condition: condition}], ruleId: ruleId, savedGroups: [{ids: [string], match: all}]}, targetId: targetId, targetType: feature-rule}], approvalNotes: approvalNotes}], templateId: templateId}",
			"--revision-comment", "revisionComment",
			"--revision-title", "revisionTitle",
			"--schedule", "{endDate: endDate, startDate: startDate}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(featuresRevisionsRulesAdd)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"features:revisions:rules", "add",
			"--id", "id",
			"--version", "new",
			"--rule", "{experimentId: experimentId, type: experiment-ref, variations: [{value: value, variationId: variationId}], allEnvironments: true, condition: condition, description: description, enabled: true, environments: [string], prerequisites: [{id: id, condition: condition}], savedGroups: [{ids: [string], match: all}], scheduleRules: [{enabled: true, timestamp: timestamp}], scheduleType: none}",
			"--ramp-schedule.end-actions", "[{patch: {condition: condition, coverage: 0, enabled: true, force: {}, prerequisites: [{id: id, condition: condition}], ruleId: ruleId, savedGroups: [{ids: [string], match: all}]}, targetId: targetId, targetType: feature-rule}]",
			"--ramp-schedule.end-condition", "{trigger: {at: {}, type: scheduled}}",
			"--ramp-schedule.name", "name",
			"--ramp-schedule.start-date", "startDate",
			"--ramp-schedule.steps", "[{trigger: {seconds: 1, type: interval}, actions: [{patch: {condition: condition, coverage: 0, enabled: true, force: {}, prerequisites: [{id: id, condition: condition}], ruleId: ruleId, savedGroups: [{ids: [string], match: all}]}, targetId: targetId, targetType: feature-rule}], approvalNotes: approvalNotes}]",
			"--ramp-schedule.template-id", "templateId",
			"--revision-comment", "revisionComment",
			"--revision-title", "revisionTitle",
			"--schedule.end-date", "endDate",
			"--schedule.start-date", "startDate",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"rule:\n" +
			"  experimentId: experimentId\n" +
			"  type: experiment-ref\n" +
			"  variations:\n" +
			"    - value: value\n" +
			"      variationId: variationId\n" +
			"  allEnvironments: true\n" +
			"  condition: condition\n" +
			"  description: description\n" +
			"  enabled: true\n" +
			"  environments:\n" +
			"    - string\n" +
			"  prerequisites:\n" +
			"    - id: id\n" +
			"      condition: condition\n" +
			"  savedGroups:\n" +
			"    - ids:\n" +
			"        - string\n" +
			"      match: all\n" +
			"  scheduleRules:\n" +
			"    - enabled: true\n" +
			"      timestamp: timestamp\n" +
			"  scheduleType: none\n" +
			"rampSchedule:\n" +
			"  endActions:\n" +
			"    - patch:\n" +
			"        condition: condition\n" +
			"        coverage: 0\n" +
			"        enabled: true\n" +
			"        force: {}\n" +
			"        prerequisites:\n" +
			"          - id: id\n" +
			"            condition: condition\n" +
			"        ruleId: ruleId\n" +
			"        savedGroups:\n" +
			"          - ids:\n" +
			"              - string\n" +
			"            match: all\n" +
			"      targetId: targetId\n" +
			"      targetType: feature-rule\n" +
			"  endCondition:\n" +
			"    trigger:\n" +
			"      at: {}\n" +
			"      type: scheduled\n" +
			"  name: name\n" +
			"  startDate: startDate\n" +
			"  steps:\n" +
			"    - trigger:\n" +
			"        seconds: 1\n" +
			"        type: interval\n" +
			"      actions:\n" +
			"        - patch:\n" +
			"            condition: condition\n" +
			"            coverage: 0\n" +
			"            enabled: true\n" +
			"            force: {}\n" +
			"            prerequisites:\n" +
			"              - id: id\n" +
			"                condition: condition\n" +
			"            ruleId: ruleId\n" +
			"            savedGroups:\n" +
			"              - ids:\n" +
			"                  - string\n" +
			"                match: all\n" +
			"          targetId: targetId\n" +
			"          targetType: feature-rule\n" +
			"      approvalNotes: approvalNotes\n" +
			"  templateId: templateId\n" +
			"revisionComment: revisionComment\n" +
			"revisionTitle: revisionTitle\n" +
			"schedule:\n" +
			"  endDate: endDate\n" +
			"  startDate: startDate\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"features:revisions:rules", "add",
			"--id", "id",
			"--version", "new",
		)
	})
}

func TestFeaturesRevisionsRulesReorder(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"features:revisions:rules", "reorder",
			"--id", "id",
			"--version", "new",
			"--rule-id", "string",
			"--revision-comment", "revisionComment",
			"--revision-title", "revisionTitle",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"ruleIds:\n" +
			"  - string\n" +
			"revisionComment: revisionComment\n" +
			"revisionTitle: revisionTitle\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"features:revisions:rules", "reorder",
			"--id", "id",
			"--version", "new",
		)
	})
}
