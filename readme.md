# Expense Tracker CLI Application

A simple command-line interface (CLI) tool for managing your personal finances. This application allows you to track expenses, delete entries, list expenses, and generate summaries by month.

https://roadmap.sh/projects/expense-tracker

## ✨ Features

- Add an expense with a description and amount.
- View a list of all expenses.
- Delete an expense by ID.
- View a summary of total expenses.
- View a summary of expenses for a specific month.
- Expenses are stored in a CSV file for easy data management.

## 📋 Requirements

- Go (version 1.17 or above)
- A terminal to run the CLI tool

## 📦 Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/yourusername/expense-tracker.git
    cd expense-tracker
    ```

2. Build the application:

    ```bash
    go build -o expense-tracker
    ```

3. Run the application:

    ```bash
    ./expense-tracker
    ```

## 🚀 Usage

### Add Expense

Add a new expense with a description and an amount:

```bash
./expense-tracker add --description "Lunch" --amount 20.50
```
Output:
```bash
Expense added successfully (ID: 1)
```
### List Expenses

View all expenses:
```bash
./expense-tracker list
```
Output:
```bash
# ID   Date        Description        Amount
# 1    2024-11-14  Lunch              $20.50
# 2    2024-11-15  Coffee             $5.00
```
### Delete Expense

Delete an expense by ID:
```bash
./expense-tracker delete --id 1
```
Output:
```bash
Expense with ID 1 deleted successfully.
```
### View Summary

View the total amount of all expenses:
```bash
./expense-tracker summary
```
Output:
```bash
Total expenses: $25.50
```
View the total expenses for a specific month:
```bash
./expense-tracker summary --month 11
```
Output:
```bash
Total expenses for month 11: $25.50
```
🗂️ Project Structure
```bash
expense-tracker/
├── main.go
├── model/
│   └── expense.go
├── services/
│   └── exp_services.go
├── results/
│   └── expense.csv
└── go.mod
```
📚 Future Enhancements

* Add categories for expenses. 
* Element monthly budget tracking. 
* Add export functionality to different file formats (e.g., JSON). 
* Add the ability to update existing expenses.

📜 License

This project is licensed under the MIT License.


### Author

Developed by Yevhenii Demenko.
Feel free to reach out on [LinkedIn](https://linkedin.com/in/demenkoeugene)