// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/growthbook/cli/internal/mocktest"
)

func TestRampSchedulesActionsAddTarget(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ramp-schedules:actions", "add-target",
			"--id", "id",
			"--feature-id", "featureId",
			"--rule-id", "ruleId",
			"--environment", "environment",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"featureId: featureId\n" +
			"ruleId: ruleId\n" +
			"environment: environment\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ramp-schedules:actions", "add-target",
			"--id", "id",
		)
	})
}

func TestRampSchedulesActionsApproveStep(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ramp-schedules:actions", "approve-step",
			"--id", "id",
		)
	})
}

func TestRampSchedulesActionsComplete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ramp-schedules:actions", "complete",
			"--id", "id",
		)
	})
}

func TestRampSchedulesActionsEjectTarget(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ramp-schedules:actions", "eject-target",
			"--id", "id",
			"--environment", "environment",
			"--rule-id", "ruleId",
			"--target-id", "targetId",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"environment: environment\n" +
			"ruleId: ruleId\n" +
			"targetId: targetId\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ramp-schedules:actions", "eject-target",
			"--id", "id",
		)
	})
}

func TestRampSchedulesActionsJump(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ramp-schedules:actions", "jump",
			"--id", "id",
			"--target-step-index", "-1",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("targetStepIndex: -1")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ramp-schedules:actions", "jump",
			"--id", "id",
		)
	})
}

func TestRampSchedulesActionsPause(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ramp-schedules:actions", "pause",
			"--id", "id",
		)
	})
}

func TestRampSchedulesActionsResume(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ramp-schedules:actions", "resume",
			"--id", "id",
		)
	})
}

func TestRampSchedulesActionsRollback(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ramp-schedules:actions", "rollback",
			"--id", "id",
		)
	})
}

func TestRampSchedulesActionsStart(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ramp-schedules:actions", "start",
			"--id", "id",
		)
	})
}
