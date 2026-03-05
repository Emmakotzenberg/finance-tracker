The fitness-tracker is a simple API to track personal expenses.

The purpose of this is to  help you keep track of your finances.

You can:
Track transactions, view lists, and get summaries.

## How to Run
1. Install Go and dependencies: `go mod tidy`
2. Run: `go run main.go`
3. Use curl for testing, e.g., `curl -X POST http://localhost:8080/transactions -d '{"amount":50,"category":"groceries"}'`
