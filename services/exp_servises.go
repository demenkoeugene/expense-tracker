package services

import (
	"errors"
	"expense-tracker/model"
	"expense-tracker/utils"
	"fmt"
	"os"
	"strconv"
	"time"
)

func DeleteExpenseByID(id int) error {
	expenses, err := utils.LoadExpanseList()
	if err != nil {
		return fmt.Errorf("error loading expenses: %w", err)
	}

	var updatedExpenses []model.Expense
	var deleted bool

	for _, expense := range expenses {
		if expense.ID == id {
			deleted = true
			continue
		}
		updatedExpenses = append(updatedExpenses, expense)
	}

	if !deleted {
		return errors.New(fmt.Sprintf("no expense found with ID: %d", id))
	}

	return utils.SaveExpenses(updatedExpenses)
}

func AddExpense(description string, amount float64) error {
	expenses, err := utils.LoadExpanseList()
	if err != nil {
		return fmt.Errorf("error loading expenses: %w", err)
	}

	for _, exp := range expenses {
		if exp.Description == description && exp.Amount == amount {
			return errors.New("this expense already exists")
		}
	}

	newExpense := model.Expense{
		ID:          len(expenses) + 1,
		Date:        time.Now(),
		Description: description,
		Amount:      amount,
	}

	expenses = append(expenses, newExpense)
	return utils.SaveExpenses(expenses)
}

func ParseOptionalMonth() int {
	if len(os.Args) > 2 && os.Args[2] == "--month" {
		if len(os.Args) < 4 {
			fmt.Println("Usage: expense-tracker summary --month <month>")
			os.Exit(1)
		}
		month, err := strconv.Atoi(os.Args[3])
		if err != nil || month < 1 || month > 12 {
			fmt.Println("Invalid month. Please enter a value between 1 and 12.")
			os.Exit(1)
		}
		return month
	}
	return 0
}
