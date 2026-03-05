## Finance Tracker
A simple personal finance tracking application built in Go, featuring a REST API, CLI commands, and a basic web dashboard for viewing summaries.

## Motivation
As a beginner developer exploring backend tools like Go and databases, I wanted to build something practical for tracking expenses. This project helps me learn about APIs, databases, concurrency basics, and integrating a simple web interface—all while creating a tool I can use daily to manage my finances without relying on complex apps.

## Goals
The primary goal of Finance Tracker is to provide an easy way to log and review expenses locally.
# Key objectives include:

1.Simple Transaction Logging: Add expenses quickly via CLI or API with minimal input.

2.Data Summaries: View totals and breakdowns by month in JSON or a web table.

3.Local Storage: Use SQLite for lightweight, file-based persistence without external services.

4.Web Dashboard: Display data in a clean HTML table for easy visualization.

5.CLI Integration: Enable terminal-based additions for fast, scriptable use.

6.Extensibility: Design for future additions like categories, budgets, or exports.

## Getting Started
Prerequisites:

*Go 1.18+ installed

## Installation

1.Clone the repo:
git clone https://github.com/Emmakotzenberg/finance-tracker.git

2.Navigate to the project:
cd finance-tracker

3.Install dependencies:
go mod tidy

## Running
Production
1.After installation, run the server:
go run main.go

2.Access the dashboard at http://localhost:8080/ to view your monthly transaction table.

3.Use CLI commands to add and remove data (see below).

4.Stop with Ctrl + C.

## Locally

The project runs entirely locally with SQLite—no external database needed.

Start the server as above.

To add transactions via CLI (without server running):

go run main.go add "AMOUNT$ DESCRIPTION DATE"

To remove transactions via CLI (without server running):
go run main.go remove "ID FROM DASHBOARD TABLE
