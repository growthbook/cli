// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/growthbook/cli/internal/mocktest"
	"github.com/growthbook/cli/internal/requestflag"
)

func TestMembersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"members", "list",
			"--global-role", "globalRole",
			"--limit", "1",
			"--offset", "0",
			"--user-email", "userEmail",
			"--user-name", "userName",
		)
	})
}

func TestMembersDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"members", "delete",
			"--id", "id",
		)
	})
}

func TestMembersUpdateRole(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"members", "update-role",
			"--id", "id",
			"--member", "{environments: [string], projectRoles: [{environments: [string], project: project, role: role, limitAccessByEnvironment: true}], role: role}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(membersUpdateRole)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"members", "update-role",
			"--id", "id",
			"--member.environments", "[string]",
			"--member.project-roles", "[{environments: [string], project: project, role: role, limitAccessByEnvironment: true}]",
			"--member.role", "role",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"member:\n" +
			"  environments:\n" +
			"    - string\n" +
			"  projectRoles:\n" +
			"    - environments:\n" +
			"        - string\n" +
			"      project: project\n" +
			"      role: role\n" +
			"      limitAccessByEnvironment: true\n" +
			"  role: role\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"members", "update-role",
			"--id", "id",
		)
	})
}
