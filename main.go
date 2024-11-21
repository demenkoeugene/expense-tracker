package main

import (
	"expense-tracker/services"
	"expense-tracker/utils"
	"fmt"
	"os"
	"strconv"
)

type CommandHandler func()

var commands = map[string]CommandHandler{
	"add":     handleAdd,
	"list":    handleList,
	"summary": handleSummary,
	"delete":  handleDelete,
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	command := os.Args[1]
	if handler, exists := commands[command]; exists {
		handler()
	} else {
		fmt.Println("Unknown command")
		printUsage()
	}
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  expense-tracker add --description <description> --amount <amount>")
	fmt.Println("  expense-tracker list")
	fmt.Println("  expense-tracker summary [--month <month>]")
	fmt.Println("  expense-tracker delete --id <id>")
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

	if err := services.DeleteExpenseByID(id); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Expense with ID %d deleted successfully.\n", id)
}

func handleSummary() {
	month := services.ParseOptionalMonth()
	expenses, err := utils.LoadExpanseList()
	if err != nil {
		fmt.Println("Error loading expenses:", err)
		return
	}

	total := utils.CalculateTotal(expenses, month)
	if month > 0 {
		fmt.Printf("Total expenses for month %d: $%.2f\n", month, total)
	} else {
		fmt.Printf("Total expenses: $%.2f\n", total)
	}
}

func handleList() {
	expenses, err := utils.LoadExpanseList()
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

	if err := services.AddExpense(description, amount); err != nil {
		fmt.Println("Error adding expense:", err)
		return
	}
	fmt.Println("Expense added successfully.")
}
