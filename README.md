# Deposit Plan Allocator (Go)

This project simulates how deposits are allocated into different investment portfolios based on a customer’s one-time and/or monthly deposit plans — similar to how StashAway works.

## 🧠 Problem Statement

A customer can:
- Create **one-time** and/or **monthly** deposit plans
- Make deposits via bank transfer (which are not linked directly to any plan)

### Goal:
Fully allocate all deposited funds across portfolios based on active plans.

## 💡 Features

- ✅ Handles one-time and monthly deposit plans
- ✅ Distributes deposits proportionally if funds are insufficient
- ✅ Written in **Go**, using integer cents for precise calculations
- ✅ Fully tested with edge cases

## 🗂 Project Structure

```
.
├── main.go               # CLI entry point (package main)
├── deposit/              # Core logic (package deposit)
│   ├── deposit.go
│   └── deposit_test.go
├── go.mod
└── go.sum
```

## 🚀 Usage

Run the CLI example:

```bash
go run main.go
```

Run the tests:

```bash
go test ./deposit
```

## 🧪 Example

Given:
- One-time plan: 
  - High risk: $10,000
  - Retirement: $500
- Monthly plan:
  - Retirement: $100

Deposits:
- $10,500
- $100

**Output:**
```
Final allocations:
High risk: $10000.00
Retirement: $600.00
```

## 📦 Requirements

- Go 1.24 or later

## 🧾 License

MIT — free to use and modify.
