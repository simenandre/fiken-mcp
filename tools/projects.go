package tools

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/simenandre/fiken-mcp-server/internal/fiken"
)

func registerProjectTools(s *server.MCPServer, client *fiken.Client) {
	s.AddTool(
		mcp.NewTool("get_projects",
			mcp.WithDescription("Returns all projects for a company (requires projects module)"),
			mcp.WithString("company_slug", mcp.Required(), mcp.Description("The company slug identifier")),
			mcp.WithNumber("page", mcp.Description("Page number (0-based)")),
			mcp.WithNumber("page_size", mcp.Description("Number of results per page (max 100)")),
			mcp.WithString("completed", mcp.Description("Filter by completed status: 'true' or 'false'")),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := req.GetArguments()
			slug := mcp.ExtractString(args, "company_slug")
			params := fiken.BuildQueryParams(
				"page", args["page"],
				"pageSize", args["page_size"],
				"completed", args["completed"],
			)
			body, status, err := client.Get("/companies/"+slug+"/projects", params)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			if status >= 400 {
				return mcp.NewToolResultError(fmt.Sprintf("API error %d: %s", status, string(body))), nil
			}
			return mcp.NewToolResultText(string(body)), nil
		},
	)

	s.AddTool(
		mcp.NewTool("get_project",
			mcp.WithDescription("Returns a specific project"),
			mcp.WithString("company_slug", mcp.Required(), mcp.Description("The company slug identifier")),
			mcp.WithString("project_id", mcp.Required(), mcp.Description("The project ID")),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := req.GetArguments()
			slug := mcp.ExtractString(args, "company_slug")
			id := mcp.ExtractString(args, "project_id")
			body, status, err := client.Get("/companies/"+slug+"/projects/"+id, nil)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			if status >= 400 {
				return mcp.NewToolResultError(fmt.Sprintf("API error %d: %s", status, string(body))), nil
			}
			return mcp.NewToolResultText(string(body)), nil
		},
	)

	s.AddTool(
		mcp.NewTool("create_project",
			mcp.WithDescription("Creates a new project"),
			mcp.WithString("company_slug", mcp.Required(), mcp.Description("The company slug identifier")),
			mcp.WithString("body", mcp.Required(), mcp.Description("JSON body with project details (name, description, startDate, endDate, etc.)")),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := req.GetArguments()
			slug := mcp.ExtractString(args, "company_slug")
			bodyStr := mcp.ExtractString(args, "body")
			body, status, err := client.Post("/companies/"+slug+"/projects", []byte(bodyStr))
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			if status >= 400 {
				return mcp.NewToolResultError(fmt.Sprintf("API error %d: %s", status, string(body))), nil
			}
			return mcp.NewToolResultText(string(body)), nil
		},
	)

	s.AddTool(
		mcp.NewTool("update_project",
			mcp.WithDescription("Updates an existing project"),
			mcp.WithString("company_slug", mcp.Required(), mcp.Description("The company slug identifier")),
			mcp.WithString("project_id", mcp.Required(), mcp.Description("The project ID")),
			mcp.WithString("body", mcp.Required(), mcp.Description("JSON body with project fields to update")),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := req.GetArguments()
			slug := mcp.ExtractString(args, "company_slug")
			id := mcp.ExtractString(args, "project_id")
			bodyStr := mcp.ExtractString(args, "body")
			body, status, err := client.Put("/companies/"+slug+"/projects/"+id, []byte(bodyStr))
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			if status >= 400 {
				return mcp.NewToolResultError(fmt.Sprintf("API error %d: %s", status, string(body))), nil
			}
			return mcp.NewToolResultText(string(body)), nil
		},
	)
}
