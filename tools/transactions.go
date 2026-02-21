package tools

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/simenandre/fiken-mcp-server/internal/fiken"
)

func registerTransactionTools(s *server.MCPServer, client *fiken.Client) {
	s.AddTool(
		mcp.NewTool("get_transactions",
			mcp.WithDescription("Returns all transactions for a company"),
			mcp.WithString("company_slug", mcp.Required(), mcp.Description("The company slug identifier")),
			mcp.WithNumber("page", mcp.Description("Page number (0-based)")),
			mcp.WithNumber("page_size", mcp.Description("Number of results per page (max 100)")),
			mcp.WithString("last_modified", mcp.Description("Filter by last modified date (YYYY-MM-DD)")),
			mcp.WithString("last_modified_le", mcp.Description("Filter: last modified ≤ date")),
			mcp.WithString("last_modified_lt", mcp.Description("Filter: last modified < date")),
			mcp.WithString("last_modified_ge", mcp.Description("Filter: last modified ≥ date")),
			mcp.WithString("last_modified_gt", mcp.Description("Filter: last modified > date")),
			mcp.WithString("created_date", mcp.Description("Filter by created date (YYYY-MM-DD)")),
			mcp.WithString("created_date_le", mcp.Description("Filter: created date ≤ value")),
			mcp.WithString("created_date_lt", mcp.Description("Filter: created date < value")),
			mcp.WithString("created_date_ge", mcp.Description("Filter: created date ≥ value")),
			mcp.WithString("created_date_gt", mcp.Description("Filter: created date > value")),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := req.GetArguments()
			slug := mcp.ExtractString(args, "company_slug")
			params := fiken.BuildQueryParams(
				"page", args["page"],
				"pageSize", args["page_size"],
				"lastModified", args["last_modified"],
				"lastModifiedLe", args["last_modified_le"],
				"lastModifiedLt", args["last_modified_lt"],
				"lastModifiedGe", args["last_modified_ge"],
				"lastModifiedGt", args["last_modified_gt"],
				"createdDate", args["created_date"],
				"createdDateLe", args["created_date_le"],
				"createdDateLt", args["created_date_lt"],
				"createdDateGe", args["created_date_ge"],
				"createdDateGt", args["created_date_gt"],
			)
			body, status, err := client.Get("/companies/"+slug+"/transactions", params)
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
