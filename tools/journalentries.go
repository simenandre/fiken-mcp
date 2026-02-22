package tools

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/simenandre/fiken-mcp/internal/fiken"
)

func registerJournalEntryTools(s *server.MCPServer, client *fiken.Client) {
	s.AddTool(
		mcp.NewTool("get_journal_entries",
			mcp.WithDescription("Returns all general journal entries for the specified company"),
			mcp.WithString("company_slug", mcp.Required(), mcp.Description("The company slug identifier")),
			mcp.WithNumber("page", mcp.Description("Page number (0-based)")),
			mcp.WithNumber("page_size", mcp.Description("Number of results per page (max 100)")),
			mcp.WithString("date", mcp.Description("Filter by date (YYYY-MM-DD)")),
			mcp.WithString("date_le", mcp.Description("Filter: date ≤ value")),
			mcp.WithString("date_lt", mcp.Description("Filter: date < value")),
			mcp.WithString("date_ge", mcp.Description("Filter: date ≥ value")),
			mcp.WithString("date_gt", mcp.Description("Filter: date > value")),
			mcp.WithString("last_modified", mcp.Description("Filter by last modified date (YYYY-MM-DD)")),
			mcp.WithString("last_modified_le", mcp.Description("Filter: last modified ≤ date")),
			mcp.WithString("last_modified_lt", mcp.Description("Filter: last modified < date")),
			mcp.WithString("last_modified_ge", mcp.Description("Filter: last modified ≥ date")),
			mcp.WithString("last_modified_gt", mcp.Description("Filter: last modified > date")),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := req.GetArguments()
			slug := mcp.ExtractString(args, "company_slug")
			params := fiken.BuildQueryParams(
				"page", args["page"],
				"pageSize", args["page_size"],
				"date", args["date"],
				"dateLe", args["date_le"],
				"dateLt", args["date_lt"],
				"dateGe", args["date_ge"],
				"dateGt", args["date_gt"],
				"lastModified", args["last_modified"],
				"lastModifiedLe", args["last_modified_le"],
				"lastModifiedLt", args["last_modified_lt"],
				"lastModifiedGe", args["last_modified_ge"],
				"lastModifiedGt", args["last_modified_gt"],
			)
			body, status, err := client.Get("/companies/"+slug+"/journalEntries", params)
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
		mcp.NewTool("get_journal_entry",
			mcp.WithDescription("Returns a specific journal entry"),
			mcp.WithString("company_slug", mcp.Required(), mcp.Description("The company slug identifier")),
			mcp.WithString("journal_entry_id", mcp.Required(), mcp.Description("The journal entry ID")),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := req.GetArguments()
			slug := mcp.ExtractString(args, "company_slug")
			id := mcp.ExtractString(args, "journal_entry_id")
			body, status, err := client.Get("/companies/"+slug+"/journalEntries/"+id, nil)
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
		mcp.NewTool("create_general_journal_entry",
			mcp.WithDescription("Creates a new general journal entry (fri postering)"),
			mcp.WithString("company_slug", mcp.Required(), mcp.Description("The company slug identifier")),
			mcp.WithString("body", mcp.Required(), mcp.Description("JSON body with journal entry details (description, date, lines, etc.)")),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := req.GetArguments()
			slug := mcp.ExtractString(args, "company_slug")
			bodyStr := mcp.ExtractString(args, "body")
			body, status, err := client.Post("/companies/"+slug+"/generalJournalEntries", []byte(bodyStr))
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
