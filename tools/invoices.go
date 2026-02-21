package tools

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/simenandre/fiken-mcp-server/internal/fiken"
)

func registerInvoiceTools(s *server.MCPServer, client *fiken.Client) {
	// Invoices
	s.AddTool(
		mcp.NewTool("get_invoices",
			mcp.WithDescription("Returns all invoices for a company"),
			mcp.WithString("company_slug", mcp.Required(), mcp.Description("The company slug identifier")),
			mcp.WithNumber("page", mcp.Description("Page number (0-based)")),
			mcp.WithNumber("page_size", mcp.Description("Number of results per page (max 100)")),
			mcp.WithString("issue_date", mcp.Description("Filter by issue date (YYYY-MM-DD)")),
			mcp.WithString("last_modified", mcp.Description("Filter by last modified date (YYYY-MM-DD)")),
			mcp.WithString("settled", mcp.Description("Filter by settled status: 'true' or 'false'")),
			mcp.WithString("customer_id", mcp.Description("Filter by customer ID")),
			mcp.WithString("order_reference", mcp.Description("Filter by order reference")),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := req.GetArguments()
			slug := mcp.ExtractString(args, "company_slug")
			params := fiken.BuildQueryParams(
				"page", args["page"],
				"pageSize", args["page_size"],
				"issueDate", args["issue_date"],
				"lastModified", args["last_modified"],
				"settled", args["settled"],
				"customerId", args["customer_id"],
				"orderReference", args["order_reference"],
			)
			body, status, err := client.Get("/companies/"+slug+"/invoices", params)
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
		mcp.NewTool("get_invoice",
			mcp.WithDescription("Returns a specific invoice"),
			mcp.WithString("company_slug", mcp.Required(), mcp.Description("The company slug identifier")),
			mcp.WithString("invoice_id", mcp.Required(), mcp.Description("The invoice ID")),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := req.GetArguments()
			slug := mcp.ExtractString(args, "company_slug")
			id := mcp.ExtractString(args, "invoice_id")
			body, status, err := client.Get("/companies/"+slug+"/invoices/"+id, nil)
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
		mcp.NewTool("create_invoice",
			mcp.WithDescription("Creates a new invoice"),
			mcp.WithString("company_slug", mcp.Required(), mcp.Description("The company slug identifier")),
			mcp.WithString("body", mcp.Required(), mcp.Description("JSON body with invoice details (issueDate, dueDate, customerId, lines, etc.)")),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := req.GetArguments()
			slug := mcp.ExtractString(args, "company_slug")
			bodyStr := mcp.ExtractString(args, "body")
			body, status, err := client.Post("/companies/"+slug+"/invoices", []byte(bodyStr))
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
		mcp.NewTool("update_invoice",
			mcp.WithDescription("Updates an existing invoice"),
			mcp.WithString("company_slug", mcp.Required(), mcp.Description("The company slug identifier")),
			mcp.WithString("invoice_id", mcp.Required(), mcp.Description("The invoice ID")),
			mcp.WithString("body", mcp.Required(), mcp.Description("JSON body with invoice fields to update")),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := req.GetArguments()
			slug := mcp.ExtractString(args, "company_slug")
			id := mcp.ExtractString(args, "invoice_id")
			bodyStr := mcp.ExtractString(args, "body")
			body, status, err := client.Put("/companies/"+slug+"/invoices/"+id, []byte(bodyStr))
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			if status >= 400 {
				return mcp.NewToolResultError(fmt.Sprintf("API error %d: %s", status, string(body))), nil
			}
			return mcp.NewToolResultText(string(body)), nil
		},
	)

	// Invoice Drafts
	s.AddTool(
		mcp.NewTool("get_invoice_drafts",
			mcp.WithDescription("Returns all invoice drafts for a company"),
			mcp.WithString("company_slug", mcp.Required(), mcp.Description("The company slug identifier")),
			mcp.WithNumber("page", mcp.Description("Page number (0-based)")),
			mcp.WithNumber("page_size", mcp.Description("Number of results per page (max 100)")),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := req.GetArguments()
			slug := mcp.ExtractString(args, "company_slug")
			params := fiken.BuildQueryParams(
				"page", args["page"],
				"pageSize", args["page_size"],
			)
			body, status, err := client.Get("/companies/"+slug+"/invoices/drafts", params)
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
		mcp.NewTool("get_invoice_draft",
			mcp.WithDescription("Returns a specific invoice draft"),
			mcp.WithString("company_slug", mcp.Required(), mcp.Description("The company slug identifier")),
			mcp.WithString("draft_id", mcp.Required(), mcp.Description("The draft ID")),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := req.GetArguments()
			slug := mcp.ExtractString(args, "company_slug")
			id := mcp.ExtractString(args, "draft_id")
			body, status, err := client.Get("/companies/"+slug+"/invoices/drafts/"+id, nil)
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
		mcp.NewTool("create_invoice_draft",
			mcp.WithDescription("Creates a new invoice draft"),
			mcp.WithString("company_slug", mcp.Required(), mcp.Description("The company slug identifier")),
			mcp.WithString("body", mcp.Required(), mcp.Description("JSON body with invoice draft details")),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := req.GetArguments()
			slug := mcp.ExtractString(args, "company_slug")
			bodyStr := mcp.ExtractString(args, "body")
			body, status, err := client.Post("/companies/"+slug+"/invoices/drafts", []byte(bodyStr))
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
		mcp.NewTool("update_invoice_draft",
			mcp.WithDescription("Updates an existing invoice draft"),
			mcp.WithString("company_slug", mcp.Required(), mcp.Description("The company slug identifier")),
			mcp.WithString("draft_id", mcp.Required(), mcp.Description("The draft ID")),
			mcp.WithString("body", mcp.Required(), mcp.Description("JSON body with invoice draft fields to update")),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := req.GetArguments()
			slug := mcp.ExtractString(args, "company_slug")
			id := mcp.ExtractString(args, "draft_id")
			bodyStr := mcp.ExtractString(args, "body")
			body, status, err := client.Put("/companies/"+slug+"/invoices/drafts/"+id, []byte(bodyStr))
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
		mcp.NewTool("delete_invoice_draft",
			mcp.WithDescription("Deletes an invoice draft"),
			mcp.WithString("company_slug", mcp.Required(), mcp.Description("The company slug identifier")),
			mcp.WithString("draft_id", mcp.Required(), mcp.Description("The draft ID")),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := req.GetArguments()
			slug := mcp.ExtractString(args, "company_slug")
			id := mcp.ExtractString(args, "draft_id")
			body, status, err := client.Delete("/companies/" + slug + "/invoices/drafts/" + id)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			if status >= 400 {
				return mcp.NewToolResultError(fmt.Sprintf("API error %d: %s", status, string(body))), nil
			}
			return mcp.NewToolResultText(fmt.Sprintf("Invoice draft %s deleted successfully", id)), nil
		},
	)

	s.AddTool(
		mcp.NewTool("create_invoice_from_draft",
			mcp.WithDescription("Creates an invoice from an existing draft"),
			mcp.WithString("company_slug", mcp.Required(), mcp.Description("The company slug identifier")),
			mcp.WithString("draft_id", mcp.Required(), mcp.Description("The draft ID")),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := req.GetArguments()
			slug := mcp.ExtractString(args, "company_slug")
			id := mcp.ExtractString(args, "draft_id")
			body, status, err := client.Post("/companies/"+slug+"/invoices/drafts/"+id+"/createInvoice", nil)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			if status >= 400 {
				return mcp.NewToolResultError(fmt.Sprintf("API error %d: %s", status, string(body))), nil
			}
			return mcp.NewToolResultText(string(body)), nil
		},
	)

	// Credit Notes
	s.AddTool(
		mcp.NewTool("get_credit_notes",
			mcp.WithDescription("Returns all credit notes for a company"),
			mcp.WithString("company_slug", mcp.Required(), mcp.Description("The company slug identifier")),
			mcp.WithNumber("page", mcp.Description("Page number (0-based)")),
			mcp.WithNumber("page_size", mcp.Description("Number of results per page (max 100)")),
			mcp.WithString("issue_date", mcp.Description("Filter by issue date (YYYY-MM-DD)")),
			mcp.WithString("last_modified", mcp.Description("Filter by last modified date (YYYY-MM-DD)")),
			mcp.WithString("settled", mcp.Description("Filter by settled status: 'true' or 'false'")),
			mcp.WithString("customer_id", mcp.Description("Filter by customer ID")),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := req.GetArguments()
			slug := mcp.ExtractString(args, "company_slug")
			params := fiken.BuildQueryParams(
				"page", args["page"],
				"pageSize", args["page_size"],
				"issueDate", args["issue_date"],
				"lastModified", args["last_modified"],
				"settled", args["settled"],
				"customerId", args["customer_id"],
			)
			body, status, err := client.Get("/companies/"+slug+"/creditNotes", params)
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
		mcp.NewTool("get_credit_note",
			mcp.WithDescription("Returns a specific credit note"),
			mcp.WithString("company_slug", mcp.Required(), mcp.Description("The company slug identifier")),
			mcp.WithString("credit_note_id", mcp.Required(), mcp.Description("The credit note ID")),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := req.GetArguments()
			slug := mcp.ExtractString(args, "company_slug")
			id := mcp.ExtractString(args, "credit_note_id")
			body, status, err := client.Get("/companies/"+slug+"/creditNotes/"+id, nil)
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
