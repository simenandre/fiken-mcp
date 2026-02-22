package tools

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/simenandre/fiken-mcp-server/internal/fiken"
)

func registerContactTools(s *server.MCPServer, client *fiken.Client) {
	s.AddTool(
		mcp.NewTool("get_contacts",
			mcp.WithDescription("Retrieves all contacts for a company"),
			mcp.WithString("company_slug", mcp.Required(), mcp.Description("The company slug identifier")),
			mcp.WithNumber("page", mcp.Description("Page number (0-based)")),
			mcp.WithNumber("page_size", mcp.Description("Number of results per page (max 100)")),
			mcp.WithString("sort_by", mcp.Description("Sort field, e.g. 'createdDate asc'")),
			mcp.WithString("name", mcp.Description("Filter by name")),
			mcp.WithString("email", mcp.Description("Filter by email")),
			mcp.WithString("organization_number", mcp.Description("Filter by organization number")),
			mcp.WithString("supplier_number", mcp.Description("Filter by supplier number")),
			mcp.WithString("customer_number", mcp.Description("Filter by customer number")),
			mcp.WithString("group", mcp.Description("Filter by group")),
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
				"sortBy", args["sort_by"],
				"name", args["name"],
				"email", args["email"],
				"organizationNumber", args["organization_number"],
				"supplierNumber", args["supplier_number"],
				"customerNumber", args["customer_number"],
				"group", args["group"],
				"lastModified", args["last_modified"],
				"lastModifiedLe", args["last_modified_le"],
				"lastModifiedLt", args["last_modified_lt"],
				"lastModifiedGe", args["last_modified_ge"],
				"lastModifiedGt", args["last_modified_gt"],
			)
			body, status, err := client.Get("/companies/"+slug+"/contacts", params)
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
		mcp.NewTool("get_contact",
			mcp.WithDescription("Retrieves a specific contact"),
			mcp.WithString("company_slug", mcp.Required(), mcp.Description("The company slug identifier")),
			mcp.WithString("contact_id", mcp.Required(), mcp.Description("The contact ID")),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := req.GetArguments()
			slug := mcp.ExtractString(args, "company_slug")
			id := mcp.ExtractString(args, "contact_id")
			body, status, err := client.Get("/companies/"+slug+"/contacts/"+id, nil)
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
		mcp.NewTool("create_contact",
			mcp.WithDescription("Creates a new contact"),
			mcp.WithString("company_slug", mcp.Required(), mcp.Description("The company slug identifier")),
			mcp.WithString("body", mcp.Required(), mcp.Description("JSON body with contact details (name, email, organizationNumber, customer, supplier, etc.)")),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := req.GetArguments()
			slug := mcp.ExtractString(args, "company_slug")
			bodyStr := mcp.ExtractString(args, "body")
			body, status, err := client.Post("/companies/"+slug+"/contacts", []byte(bodyStr))
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
		mcp.NewTool("update_contact",
			mcp.WithDescription("Updates an existing contact"),
			mcp.WithString("company_slug", mcp.Required(), mcp.Description("The company slug identifier")),
			mcp.WithString("contact_id", mcp.Required(), mcp.Description("The contact ID")),
			mcp.WithString("body", mcp.Required(), mcp.Description("JSON body with contact fields to update")),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := req.GetArguments()
			slug := mcp.ExtractString(args, "company_slug")
			id := mcp.ExtractString(args, "contact_id")
			bodyStr := mcp.ExtractString(args, "body")
			body, status, err := client.Put("/companies/"+slug+"/contacts/"+id, []byte(bodyStr))
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
		mcp.NewTool("delete_contact",
			mcp.WithDescription("Deletes a contact (or sets to inactive if it has associated records)"),
			mcp.WithString("company_slug", mcp.Required(), mcp.Description("The company slug identifier")),
			mcp.WithString("contact_id", mcp.Required(), mcp.Description("The contact ID")),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := req.GetArguments()
			slug := mcp.ExtractString(args, "company_slug")
			id := mcp.ExtractString(args, "contact_id")
			body, status, err := client.Delete("/companies/" + slug + "/contacts/" + id)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			if status >= 400 {
				return mcp.NewToolResultError(fmt.Sprintf("API error %d: %s", status, string(body))), nil
			}
			if len(body) > 0 {
				return mcp.NewToolResultText(string(body)), nil
			}
			return mcp.NewToolResultText(fmt.Sprintf("Contact %s deleted successfully", id)), nil
		},
	)
}
