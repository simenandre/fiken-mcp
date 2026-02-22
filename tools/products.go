package tools

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/simenandre/fiken-mcp/internal/fiken"
)

func registerProductTools(s *server.MCPServer, client *fiken.Client) {
	s.AddTool(
		mcp.NewTool("get_products",
			mcp.WithDescription("Returns all products for a company"),
			mcp.WithString("company_slug", mcp.Required(), mcp.Description("The company slug identifier")),
			mcp.WithNumber("page", mcp.Description("Page number (0-based)")),
			mcp.WithNumber("page_size", mcp.Description("Number of results per page (max 100)")),
			mcp.WithString("name", mcp.Description("Filter by product name")),
			mcp.WithString("product_number", mcp.Description("Filter by product number")),
			mcp.WithString("active", mcp.Description("Filter by active status: 'true' or 'false'")),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := req.GetArguments()
			slug := mcp.ExtractString(args, "company_slug")
			params := fiken.BuildQueryParams(
				"page", args["page"],
				"pageSize", args["page_size"],
				"name", args["name"],
				"productNumber", args["product_number"],
				"active", args["active"],
			)
			body, status, err := client.Get("/companies/"+slug+"/products", params)
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
		mcp.NewTool("get_product",
			mcp.WithDescription("Returns a specific product"),
			mcp.WithString("company_slug", mcp.Required(), mcp.Description("The company slug identifier")),
			mcp.WithString("product_id", mcp.Required(), mcp.Description("The product ID")),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := req.GetArguments()
			slug := mcp.ExtractString(args, "company_slug")
			id := mcp.ExtractString(args, "product_id")
			body, status, err := client.Get("/companies/"+slug+"/products/"+id, nil)
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
		mcp.NewTool("create_product",
			mcp.WithDescription("Creates a new product"),
			mcp.WithString("company_slug", mcp.Required(), mcp.Description("The company slug identifier")),
			mcp.WithString("body", mcp.Required(), mcp.Description("JSON body with product details (name, unitPrice, incomeAccount, vatType, etc.)")),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := req.GetArguments()
			slug := mcp.ExtractString(args, "company_slug")
			bodyStr := mcp.ExtractString(args, "body")
			body, status, err := client.Post("/companies/"+slug+"/products", []byte(bodyStr))
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
		mcp.NewTool("update_product",
			mcp.WithDescription("Updates an existing product"),
			mcp.WithString("company_slug", mcp.Required(), mcp.Description("The company slug identifier")),
			mcp.WithString("product_id", mcp.Required(), mcp.Description("The product ID")),
			mcp.WithString("body", mcp.Required(), mcp.Description("JSON body with product fields to update")),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := req.GetArguments()
			slug := mcp.ExtractString(args, "company_slug")
			id := mcp.ExtractString(args, "product_id")
			bodyStr := mcp.ExtractString(args, "body")
			body, status, err := client.Put("/companies/"+slug+"/products/"+id, []byte(bodyStr))
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
		mcp.NewTool("delete_product",
			mcp.WithDescription("Deletes a product"),
			mcp.WithString("company_slug", mcp.Required(), mcp.Description("The company slug identifier")),
			mcp.WithString("product_id", mcp.Required(), mcp.Description("The product ID")),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := req.GetArguments()
			slug := mcp.ExtractString(args, "company_slug")
			id := mcp.ExtractString(args, "product_id")
			body, status, err := client.Delete("/companies/" + slug + "/products/" + id)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			if status >= 400 {
				return mcp.NewToolResultError(fmt.Sprintf("API error %d: %s", status, string(body))), nil
			}
			return mcp.NewToolResultText(fmt.Sprintf("Product %s deleted successfully", id)), nil
		},
	)
}
