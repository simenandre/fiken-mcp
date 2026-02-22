package tools

import (
	"github.com/mark3labs/mcp-go/server"
	"github.com/simenandre/fiken-mcp/internal/fiken"
)

// RegisterAll registers all Fiken tools with the MCP server.
func RegisterAll(s *server.MCPServer, client *fiken.Client) {
	registerUserTools(s, client)
	registerCompanyTools(s, client)
	registerAccountTools(s, client)
	registerBankAccountTools(s, client)
	registerContactTools(s, client)
	registerJournalEntryTools(s, client)
	registerTransactionTools(s, client)
	registerProductTools(s, client)
	registerInvoiceTools(s, client)
	registerPurchaseTools(s, client)
	registerSalesTools(s, client)
	registerProjectTools(s, client)
	registerOfferTools(s, client)
	registerOrderConfirmationTools(s, client)
	registerInboxTools(s, client)
}
