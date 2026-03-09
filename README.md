## Finance Tracker

A simple personal finance tracking tool built in Go.
It allows users to add, remove, and view transactions via a web dashboard or CLI, using a local SQLite database.

## Motivation

This project was created to provide an easy, privacy-focused way to track monthly transactions without relying on cloud services or external databases.
It's ideal for individuals who want a lightweight tool to monitor expenses, income, and financial patterns directly on their local machine.

## Quick Start
# Prerequisites
Go 1.18+ installed

# Installation
1.Clone the repo
git clone https://github.com/Emmakotzenberg/finance-tracker.git

2.Navigate to the project
"cd finance-tracker"

3.Install dependencies
"go mod tidy"

## Running the Server
1.After installation, run the server:
"go run main.go"

2.Access the dashboard at http://localhost:8080/ to view your monthly transactions.

3.Use CLI commands to add and remove data (see Usage section below).

4.Stop the server with Ctrl + C.

## Usage
# Adding Transactions via CLI
"go run main.go add "AMOUNT$ DESCRIPTION DATE"

-Replace AMOUNT$ with your transaction amount
-DESCRIPTION is a brief note of the transaction
-DATE should be YYYY-MM-DD format

# Removing Transactions via CLI 
"go run main.go remove "ID"

-Replace ID with transaction ID visible in web dashboard

## Contributing

Contributions are welcome! To get started:

1.Fork the repository on GitHub.
 
2.Create a new branch for your feature or bug fix: "git checkout -b feature/your-feature-name".

3.Make your changes and commit them with descriptive messages.

4.Push your branch: "git push origin feature/your-feature-name".

5.Open a Pull Request on the original repository.

Please follow the code style in the project, add tests where possible, and ensure your changes don't break existing functionality.
For major changes, open an issue first to discuss.

