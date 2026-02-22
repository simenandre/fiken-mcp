# fiken-mcp-server

A [Model Context Protocol (MCP)](https://modelcontextprotocol.io) server for the [Fiken](https://fiken.no) accounting API. Connect your AI assistant (Claude, Copilot, etc.) to your Fiken account to read and manage accounting data using natural language.

## Prerequisites

- [Go](https://go.dev/) 1.24 or later
- A Fiken account with an [API token](https://fiken.no/api/v2/documentation/#section/Authentication)

## Installation

Install directly with `go install`:

```sh
go install github.com/simenandre/fiken-mcp@latest
```

Or clone and build from source:

```sh
git clone https://github.com/simenandre/fiken-mcp.git
cd fiken-mcp
go build -o fiken-mcp .
```

## Configuration

The server requires a Fiken API token set via the `FIKEN_API_KEY` environment variable.

You can generate an API token in your Fiken account under **Settings → API**.

### Claude Desktop

Add the following to your `claude_desktop_config.json`:

```json
{
  "mcpServers": {
    "fiken": {
      "command": "fiken-mcp",
      "env": {
        "FIKEN_API_KEY": "your-api-token-here"
      }
    }
  }
}
```

The config file is located at:
- **macOS**: `~/Library/Application Support/Claude/claude_desktop_config.json`
- **Windows**: `%APPDATA%\Claude\claude_desktop_config.json`

### VS Code (GitHub Copilot)

Add the following to your `.vscode/mcp.json` or user settings:

```json
{
  "servers": {
    "fiken": {
      "type": "stdio",
      "command": "fiken-mcp",
      "env": {
        "FIKEN_API_KEY": "your-api-token-here"
      }
    }
  }
}
```

## Available Tools

The server exposes the following tools to your AI assistant:

### User
| Tool | Description |
|------|-------------|
| `get_user` | Get information about the authenticated user |

### Companies
| Tool | Description |
|------|-------------|
| `get_companies` | List all companies the user has access to |
| `get_company` | Get details for a specific company |

### Accounts
| Tool | Description |
|------|-------------|
| `get_accounts` | List bookkeeping accounts |
| `get_account` | Get a specific bookkeeping account |
| `get_account_balances` | Get account balances for a given date |
| `get_account_balance` | Get the balance for a specific account |

### Bank Accounts
| Tool | Description |
|------|-------------|
| `get_bank_accounts` | List all bank accounts |
| `get_bank_account` | Get a specific bank account |
| `create_bank_account` | Create a new bank account |
| `get_bank_balances` | Get bank balances |

### Contacts
| Tool | Description |
|------|-------------|
| `get_contacts` | List contacts |
| `get_contact` | Get a specific contact |
| `create_contact` | Create a new contact |
| `update_contact` | Update an existing contact |
| `delete_contact` | Delete a contact |

### Invoices
| Tool | Description |
|------|-------------|
| `get_invoices` | List invoices |
| `get_invoice` | Get a specific invoice |
| `create_invoice` | Create a new invoice |
| `update_invoice` | Update an invoice |
| `get_invoice_drafts` | List invoice drafts |
| `get_invoice_draft` | Get a specific invoice draft |
| `create_invoice_draft` | Create an invoice draft |
| `update_invoice_draft` | Update an invoice draft |
| `delete_invoice_draft` | Delete an invoice draft |
| `create_invoice_from_draft` | Create an invoice from a draft |
| `get_credit_notes` | List credit notes |
| `get_credit_note` | Get a specific credit note |

### Journal Entries
| Tool | Description |
|------|-------------|
| `get_journal_entries` | List general journal entries |
| `get_journal_entry` | Get a specific journal entry |
| `create_general_journal_entry` | Create a new general journal entry |

### Transactions
| Tool | Description |
|------|-------------|
| `get_transactions` | List all transactions |

### Products
| Tool | Description |
|------|-------------|
| `get_products` | List products |
| `get_product` | Get a specific product |
| `create_product` | Create a new product |
| `update_product` | Update a product |
| `delete_product` | Delete a product |

### Purchases
| Tool | Description |
|------|-------------|
| `get_purchases` | List purchases |
| `get_purchase` | Get a specific purchase |
| `create_purchase` | Create a new purchase |
| `get_purchase_drafts` | List purchase drafts |
| `get_purchase_draft` | Get a specific purchase draft |
| `create_purchase_draft` | Create a purchase draft |
| `delete_purchase_draft` | Delete a purchase draft |
| `create_purchase_from_draft` | Create a purchase from a draft |

### Sales
| Tool | Description |
|------|-------------|
| `get_sales` | List sales |
| `get_sale` | Get a specific sale |
| `create_sale` | Create a new sale |
| `get_sale_drafts` | List sale drafts |
| `get_sale_draft` | Get a specific sale draft |
| `create_sale_draft` | Create a sale draft |
| `delete_sale_draft` | Delete a sale draft |
| `create_sale_from_draft` | Create a sale from a draft |

### Projects
| Tool | Description |
|------|-------------|
| `get_projects` | List projects |
| `get_project` | Get a specific project |
| `create_project` | Create a new project |
| `update_project` | Update a project |

### Offers
| Tool | Description |
|------|-------------|
| `get_offers` | List all offers |
| `get_offer` | Get a specific offer |

### Order Confirmations
| Tool | Description |
|------|-------------|
| `get_order_confirmations` | List all order confirmations |
| `get_order_confirmation` | Get a specific order confirmation |

### Inbox
| Tool | Description |
|------|-------------|
| `get_inbox` | List documents in the inbox |
| `get_inbox_item` | Get a specific inbox document |

## Development

Run unit tests:

```sh
go test ./...
```

Run integration tests (requires a valid `FIKEN_API_KEY`):

```sh
go test -tags integration ./...
```
