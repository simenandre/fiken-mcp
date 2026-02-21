package tools

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/simenandre/fiken-mcp-server/internal/fiken"
)

func registerAccountTools(s *server.MCPServer, client *fiken.Client) {
	s.AddTool(
		mcp.NewTool("get_accounts",
			mcp.WithDescription("Retrieves the bookkeeping accounts for the current year"),
			mcp.WithString("company_slug", mcp.Required(), mcp.Description("The company slug identifier")),
			mcp.WithString("from_account", mcp.Description("Filter: from account number")),
			mcp.WithString("to_account", mcp.Description("Filter: to account number")),
			mcp.WithNumber("page", mcp.Description("Page number (0-based)")),
			mcp.WithNumber("page_size", mcp.Description("Number of results per page (max 100)")),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := req.GetArguments()
			slug := mcp.ExtractString(args, "company_slug")
			params := fiken.BuildQueryParams(
				"fromAccount", args["from_account"],
				"toAccount", args["to_account"],
				"page", args["page"],
				"pageSize", args["page_size"],
			)
			body, status, err := client.Get("/companies/"+slug+"/accounts", params)
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
		mcp.NewTool("get_account",
			mcp.WithDescription("Retrieves a specific bookkeeping account"),
			mcp.WithString("company_slug", mcp.Required(), mcp.Description("The company slug identifier")),
			mcp.WithString("account_code", mcp.Required(), mcp.Description("The account code (e.g. '3020' or '1500:10001')")),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := req.GetArguments()
			slug := mcp.ExtractString(args, "company_slug")
			code := mcp.ExtractString(args, "account_code")
			body, status, err := client.Get("/companies/"+slug+"/accounts/"+code, nil)
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
		mcp.NewTool("get_account_balances",
			mcp.WithDescription("Retrieves bookkeeping accounts and closing balances for a given date"),
			mcp.WithString("company_slug", mcp.Required(), mcp.Description("The company slug identifier")),
			mcp.WithString("date", mcp.Required(), mcp.Description("Date in YYYY-MM-DD format")),
			mcp.WithString("from_account", mcp.Description("Filter: from account number")),
			mcp.WithString("to_account", mcp.Description("Filter: to account number")),
			mcp.WithNumber("page", mcp.Description("Page number (0-based)")),
			mcp.WithNumber("page_size", mcp.Description("Number of results per page (max 100)")),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := req.GetArguments()
			slug := mcp.ExtractString(args, "company_slug")
			params := fiken.BuildQueryParams(
				"date", args["date"],
				"fromAccount", args["from_account"],
				"toAccount", args["to_account"],
				"page", args["page"],
				"pageSize", args["page_size"],
			)
			body, status, err := client.Get("/companies/"+slug+"/accountBalances", params)
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
		mcp.NewTool("get_account_balance",
			mcp.WithDescription("Retrieves balance for a specific bookkeeping account on a given date"),
			mcp.WithString("company_slug", mcp.Required(), mcp.Description("The company slug identifier")),
			mcp.WithString("account_code", mcp.Required(), mcp.Description("The account code")),
			mcp.WithString("date", mcp.Required(), mcp.Description("Date in YYYY-MM-DD format")),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := req.GetArguments()
			slug := mcp.ExtractString(args, "company_slug")
			code := mcp.ExtractString(args, "account_code")
			params := fiken.BuildQueryParams("date", args["date"])
			body, status, err := client.Get("/companies/"+slug+"/accountBalances/"+code, params)
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
