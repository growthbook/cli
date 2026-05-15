// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/growthbook-cli/internal/mocktest"
	"github.com/stainless-sdks/growthbook-cli/internal/requestflag"
)

func TestTeamsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"teams", "create",
			"--description", "description",
			"--name", "name",
			"--role", "role",
			"--created-by", "createdBy",
			"--default-project", "defaultProject",
			"--environment", "string",
			"--limit-access-by-environment=true",
			"--managed-by", "{resourceId: resourceId, type: vercel}",
			"--project-role", "{environments: [string], limitAccessByEnvironment: true, project: project, role: role, teams: [string]}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(teamsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"teams", "create",
			"--description", "description",
			"--name", "name",
			"--role", "role",
			"--created-by", "createdBy",
			"--default-project", "defaultProject",
			"--environment", "string",
			"--limit-access-by-environment=true",
			"--managed-by.resource-id", "resourceId",
			"--managed-by.type", "vercel",
			"--project-role.environments", "[string]",
			"--project-role.limit-access-by-environment=true",
			"--project-role.project", "project",
			"--project-role.role", "role",
			"--project-role.teams", "[string]",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"description: description\n" +
			"name: name\n" +
			"role: role\n" +
			"createdBy: createdBy\n" +
			"defaultProject: defaultProject\n" +
			"environments:\n" +
			"  - string\n" +
			"limitAccessByEnvironment: true\n" +
			"managedBy:\n" +
			"  resourceId: resourceId\n" +
			"  type: vercel\n" +
			"projectRoles:\n" +
			"  - environments:\n" +
			"      - string\n" +
			"    limitAccessByEnvironment: true\n" +
			"    project: project\n" +
			"    role: role\n" +
			"    teams:\n" +
			"      - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"teams", "create",
		)
	})
}

func TestTeamsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"teams", "retrieve",
			"--id", "id",
		)
	})
}

func TestTeamsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"teams", "update",
			"--id", "id",
			"--created-by", "createdBy",
			"--default-project", "defaultProject",
			"--description", "description",
			"--environment", "string",
			"--limit-access-by-environment=true",
			"--managed-by", "{resourceId: resourceId, type: vercel}",
			"--name", "name",
			"--project-role", "{environments: [string], limitAccessByEnvironment: true, project: project, role: role, teams: [string]}",
			"--role", "role",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(teamsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"teams", "update",
			"--id", "id",
			"--created-by", "createdBy",
			"--default-project", "defaultProject",
			"--description", "description",
			"--environment", "string",
			"--limit-access-by-environment=true",
			"--managed-by.resource-id", "resourceId",
			"--managed-by.type", "vercel",
			"--name", "name",
			"--project-role.environments", "[string]",
			"--project-role.limit-access-by-environment=true",
			"--project-role.project", "project",
			"--project-role.role", "role",
			"--project-role.teams", "[string]",
			"--role", "role",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"createdBy: createdBy\n" +
			"defaultProject: defaultProject\n" +
			"description: description\n" +
			"environments:\n" +
			"  - string\n" +
			"limitAccessByEnvironment: true\n" +
			"managedBy:\n" +
			"  resourceId: resourceId\n" +
			"  type: vercel\n" +
			"name: name\n" +
			"projectRoles:\n" +
			"  - environments:\n" +
			"      - string\n" +
			"    limitAccessByEnvironment: true\n" +
			"    project: project\n" +
			"    role: role\n" +
			"    teams:\n" +
			"      - string\n" +
			"role: role\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"teams", "update",
			"--id", "id",
		)
	})
}

func TestTeamsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"teams", "list",
		)
	})
}

func TestTeamsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"teams", "delete",
			"--id", "id",
			"--delete-members", "deleteMembers",
		)
	})
}
