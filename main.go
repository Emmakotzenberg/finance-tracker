package main

import (
    "database/sql"
    "net/http"
    "time"
    "sort"

    "github.com/gin-gonic/gin"
    _ "github.com/mattn/go-sqlite3"
)

type Transaction struct {
    ID          int       `json:"id"`
    Amount      float64   `json:"amount"`
    Category    string    `json:"category"`
    Date        time.Time `json:"date"`
    Description string    `json:"description"`
}

type MonthlySummary struct {
    Month        string        `json:"month"`
    Transactions []Transaction `json:"transactions"`
    Total        float64       `json:"total"`
}

var db *sql.DB

func main() {
    var err error
    db, err = sql.Open("sqlite3", "./tracker.db")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    // Create table if not exists
    _, err = db.Exec(`CREATE TABLE IF NOT EXISTS transactions (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        amount REAL,
        category TEXT,
        date TEXT,
        description TEXT
    )`)
    if err != nil {
        panic(err)
    }

    r := gin.Default()
    r.POST("/transactions", addTransaction)
    r.GET("/transactions", listTransactions)
    r.GET("/summary", summary)
    r.GET("/monthly-summary", getMonthlySummary)
    r.Run(":8080")
}

func addTransaction(c *gin.Context) {
    var tx Transaction
    if err := c.BindJSON(&tx); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    tx.Date = time.Now() // Or parse from input
    _, err := db.Exec("INSERT INTO transactions (amount, category, date, description) VALUES (?, ?, ?, ?)",
        tx.Amount, tx.Category, tx.Date.Format(time.RFC3339), tx.Description)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Transaction added"})
}

func listTransactions(c *gin.Context) {
    rows, err := db.Query("SELECT * FROM transactions")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()

    var transactions []Transaction
    for rows.Next() {
        var tx Transaction
        var dateStr string
        if err := rows.Scan(&tx.ID, &tx.Amount, &tx.Category, &dateStr, &tx.Description); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        tx.Date, _ = time.Parse(time.RFC3339, dateStr)
        transactions = append(transactions, tx)
    }
    c.JSON(http.StatusOK, transactions)
}

func summary(c *gin.Context) {
    // Simple total spent
    var total float64
    err := db.QueryRow("SELECT SUM(amount) FROM transactions").Scan(&total)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"total_spent": total})
}

func getMonthlySummary(c *gin.Context) {
    rows, err := db.Query("SELECT * FROM transactions ORDER BY date DESC")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()

    monthly := make(map[string][]Transaction)
    totals := make(map[string]float64)

    for rows.Next() {
        var tx Transaction
        var dateStr string
        if err := rows.Scan(&tx.ID, &tx.Amount, &tx.Category, &dateStr, &tx.Description); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        tx.Date, _ = time.Parse(time.RFC3339, dateStr)

        monthKey := tx.Date.Format("2006-01")           // e.g. "2026-03"
        monthly[monthKey] = append(monthly[monthKey], tx)
        totals[monthKey] += tx.Amount
    }

    var summaries []MonthlySummary
    for month, txs := range monthly {
        summaries = append(summaries, MonthlySummary{
            Month:        month,
            Transactions: txs,
            Total:        totals[month],
        })
    }

    // Show newest month first
    sort.Slice(summaries, func(i, j int) bool {
        return summaries[i].Month > summaries[j].Month
    })

    c.JSON(http.StatusOK, summaries)
}
