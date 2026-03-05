package main

import (
    "database/sql"
    "fmt"
    "net/http"
    "time"

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
