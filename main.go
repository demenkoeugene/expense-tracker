package main

import (
	"expense-tracker/model"
	"expense-tracker/services"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Print("Usage: expense-tracker <command> [arguments]")
		return
	}

	command := os.Args[1]
	switch command {
	case "add":
		handleAdd()
	case "list":
		handleList()
	case "summary":
		handleSummary()
	case "delete":
		handleDelete()
	default:
		fmt.Println("Unknown command")
	}
}

func handleDelete() {
	if len(os.Args) < 4 || os.Args[2] != "--id" {
		fmt.Println("Usage: expense-tracker delete --id <id>")
		return
	}

	id, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("Invalid ID. Please enter a valid number.")
		return
	}

	expenses, err := services.LoadExpanseList()
	if err != nil {
		fmt.Println("Error loading expenses:", err)
		return
	}

	var updatedExpenses []model.Expense
	var deleted bool

	for _, expense := range expenses {
		if expense.ID == id {
			fmt.Printf("Deleting expense (ID: %d)\n", expense.ID)
			deleted = true
			continue
		}
		updatedExpenses = append(updatedExpenses, expense)
	}

	if !deleted {
		fmt.Printf("No expense found with ID: %d\n", id)
		return
	}

	err = services.SaveExpenses(updatedExpenses)
	if err != nil {
		fmt.Println("Error saving updated expenses:", err)
		return
	}

	fmt.Printf("Expense with ID %d deleted successfully.\n", id)
}

func handleSummary() {
	var month int
	if len(os.Args) > 2 && os.Args[2] == "--month" {
		if len(os.Args) < 4 {
			fmt.Println("Usage: expense-tracker summary --month <month>")
			return
		}
		var err error
		month, err = strconv.Atoi(os.Args[3])
		if err != nil || month < 1 || month > 12 {
			fmt.Println("Invalid month. Please enter a value between 1 and 12.")
			return
		}
	}

	expenses, err := services.LoadExpanseList()
	if err != nil {
		fmt.Println("Error loading expenses:", err)
		return
	}

	total := services.CalculateTotal(expenses, month)
	if month > 0 {
		fmt.Printf("Total expenses for month %d: $%.2f\n", month, total)
	} else {
		fmt.Printf("Total expenses: $%.2f\n", total)
	}
}

func handleList() {
	expenses, err := services.LoadExpanseList()
	if err != nil {
		fmt.Println("Error loading expenses:", err)
		return
	}

	fmt.Printf("# %-4s %-12s %-20s %s\n", "ID", "Date", "Description", "Amount")
	for _, exp := range expenses {
		exp.Print()
	}
}

func handleAdd() {
	if len(os.Args) < 6 || os.Args[2] != "--description" || os.Args[4] != "--amount" {
		fmt.Println("Usage: expense-tracker add --description <description> --amount <amount>")
		return
	}

	description := os.Args[3]
	amount, err := strconv.ParseFloat(os.Args[5], 64)
	if err != nil {
		fmt.Println("Invalid amount. Please enter a valid number.")
		return
	}

	expenses, err := services.LoadExpanseList()
	if err != nil {
		fmt.Println("Error loading expenses:", err)
		return
	}

	// Перевірка на дублікат
	for _, exp := range expenses {
		if exp.Description == description && exp.Amount == amount {
			fmt.Println("This expense already exists.")
			return
		}
	}

	newExpense := model.Expense{
		ID:          len(expenses) + 1,
		Date:        time.Now(),
		Description: description,
		Amount:      amount,
	}

	expenses = append(expenses, newExpense)
	err = services.SaveExpenses(expenses)
	if err != nil {
		fmt.Println("Error saving expenses:", err)
		return
	}
	fmt.Printf("Expense added successfully (ID: %d)\n", newExpense.ID)
}
