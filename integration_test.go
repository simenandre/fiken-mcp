//go:build integration

package main

import (
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/simenandre/fiken-mcp/internal/fiken"
)

const testCompanySlug = "fiken-demo-lastebil-og-slott-as1"

// demoAPIKey is the publicly shared sandbox demo key provided in the project requirements
// for the fiken-demo-lastebil-og-slott-as1 test company. It is not a production credential.
// In CI or production environments, set the FIKEN_API_KEY environment variable instead.
// WARNING: Never use a real production API key in source code.
const demoAPIKey = "11594281821.4UARszG4G0yPqvUZmJ0JT6C1ltVQYAUm"

func newTestClient() *fiken.Client {
	apiKey := os.Getenv("FIKEN_API_KEY")
	if apiKey == "" {
		apiKey = demoAPIKey
	}
	return fiken.NewClient(apiKey)
}

func truncate(b []byte, n int) string {
	s := string(b)
	if len(s) > n {
		return s[:n]
	}
	return s
}

func TestGetUser(t *testing.T) {
	client := newTestClient()
	body, status, err := client.Get("/user", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if status != 200 {
		t.Fatalf("expected status 200, got %d: %s", status, string(body))
	}
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		t.Fatalf("failed to parse response: %v", err)
	}
	if _, ok := result["name"]; !ok {
		t.Error("expected 'name' field in user response")
	}
	t.Logf("User: %s", string(body))
}

func TestGetCompanies(t *testing.T) {
	client := newTestClient()
	body, status, err := client.Get("/companies", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if status != 200 {
		t.Fatalf("expected status 200, got %d: %s", status, string(body))
	}
	var result []interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		t.Fatalf("failed to parse response: %v", err)
	}
	if len(result) == 0 {
		t.Error("expected at least one company")
	}
	t.Logf("Found %d companies", len(result))
}

func TestGetCompany(t *testing.T) {
	client := newTestClient()
	body, status, err := client.Get("/companies/"+testCompanySlug, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if status != 200 {
		t.Fatalf("expected status 200, got %d: %s", status, string(body))
	}
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		t.Fatalf("failed to parse response: %v", err)
	}
	if result["slug"] != testCompanySlug {
		t.Errorf("expected slug %s, got %v", testCompanySlug, result["slug"])
	}
	t.Logf("Company: %s", result["name"])
}

func TestGetContacts(t *testing.T) {
	client := newTestClient()
	body, status, err := client.Get("/companies/"+testCompanySlug+"/contacts", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if status != 200 {
		t.Fatalf("expected status 200, got %d: %s", status, string(body))
	}
	t.Logf("Contacts: %s", truncate(body, 200))
}

func TestCreateAndGetContact(t *testing.T) {
	client := newTestClient()

	createBody := []byte(`{"name":"Test Contact MCP","email":"testmcp@example.com","customer":true,"supplier":false}`)
	respBody, status, err := client.Post("/companies/"+testCompanySlug+"/contacts", createBody)
	if err != nil {
		t.Fatalf("unexpected error creating contact: %v", err)
	}
	if status != 201 {
		t.Fatalf("expected status 201, got %d: %s", status, string(respBody))
	}
	t.Logf("Created contact at: %s", string(respBody))

	// Verify by listing contacts filtered by name
	listBody, status2, err := client.Get("/companies/"+testCompanySlug+"/contacts", map[string]string{"name": "Test Contact MCP"})
	if err != nil {
		t.Fatalf("unexpected error listing contacts: %v", err)
	}
	if status2 != 200 {
		t.Fatalf("expected status 200, got %d: %s", status2, string(listBody))
	}
	if !strings.Contains(string(listBody), "Test Contact MCP") {
		t.Error("created contact not found in list")
	}
	t.Logf("Found created contact in list")
}

func TestGetAccounts(t *testing.T) {
	client := newTestClient()
	body, status, err := client.Get("/companies/"+testCompanySlug+"/accounts", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if status != 200 {
		t.Fatalf("expected status 200, got %d: %s", status, string(body))
	}
	t.Logf("Accounts: %s", truncate(body, 200))
}

func TestGetProducts(t *testing.T) {
	client := newTestClient()
	body, status, err := client.Get("/companies/"+testCompanySlug+"/products", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if status != 200 {
		t.Fatalf("expected status 200, got %d: %s", status, string(body))
	}
	t.Logf("Products: %s", truncate(body, 200))
}

func TestGetInvoices(t *testing.T) {
	client := newTestClient()
	body, status, err := client.Get("/companies/"+testCompanySlug+"/invoices", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if status != 200 {
		t.Fatalf("expected status 200, got %d: %s", status, string(body))
	}
	t.Logf("Invoices: %s", truncate(body, 200))
}

func TestGetBankAccounts(t *testing.T) {
	client := newTestClient()
	body, status, err := client.Get("/companies/"+testCompanySlug+"/bankAccounts", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if status != 200 {
		t.Fatalf("expected status 200, got %d: %s", status, string(body))
	}
	t.Logf("Bank accounts: %s", truncate(body, 200))
}

func TestGetJournalEntries(t *testing.T) {
	client := newTestClient()
	body, status, err := client.Get("/companies/"+testCompanySlug+"/journalEntries", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if status != 200 {
		t.Fatalf("expected status 200, got %d: %s", status, string(body))
	}
	t.Logf("Journal entries: %s", truncate(body, 200))
}

func TestGetTransactions(t *testing.T) {
	client := newTestClient()
	body, status, err := client.Get("/companies/"+testCompanySlug+"/transactions", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if status != 200 {
		t.Fatalf("expected status 200, got %d: %s", status, string(body))
	}
	t.Logf("Transactions: %s", truncate(body, 200))
}

func TestGetSales(t *testing.T) {
	client := newTestClient()
	body, status, err := client.Get("/companies/"+testCompanySlug+"/sales", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if status != 200 {
		t.Fatalf("expected status 200, got %d: %s", status, string(body))
	}
	t.Logf("Sales: %s", truncate(body, 200))
}

func TestGetPurchases(t *testing.T) {
	client := newTestClient()
	body, status, err := client.Get("/companies/"+testCompanySlug+"/purchases", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if status != 200 {
		t.Fatalf("expected status 200, got %d: %s", status, string(body))
	}
	t.Logf("Purchases: %s", truncate(body, 200))
}

func TestGetProjects(t *testing.T) {
	client := newTestClient()
	body, status, err := client.Get("/companies/"+testCompanySlug+"/projects", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// 402 means the projects/time-tracking module is not activated for this demo account
	if status == 402 {
		t.Skipf("Projects module not activated (402): %s", string(body))
	}
	if status != 200 {
		t.Fatalf("expected status 200, got %d: %s", status, string(body))
	}
	t.Logf("Projects: %s", truncate(body, 200))
}
