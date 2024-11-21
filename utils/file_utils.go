package utils

import (
	"encoding/csv"
	"expense-tracker/model"
	"fmt"
	"os"
	"strconv"
	"time"
)

const expanseFile = "results/expense.csv"

func LoadExpanseList() ([]model.Expense, error) {
	file, err := os.Open(expanseFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []model.Expense{}, nil
		}
		return nil, err
	}
	defer file.Close()

	return parseCSV(file)
}

func parseCSV(file *os.File) ([]model.Expense, error) {
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var expenses []model.Expense
	for _, record := range records {
		expense, err := parseRecord(record)
		if err != nil {
			return nil, err
		}
		expenses = append(expenses, expense)
	}
	return expenses, nil
}

func parseRecord(record []string) (model.Expense, error) {
	id, err := strconv.Atoi(record[0])
	if err != nil {
		return model.Expense{}, fmt.Errorf("invalid ID: %v", err)
	}

	date, err := time.Parse("2006-01-02", record[1])
	if err != nil {
		return model.Expense{}, fmt.Errorf("invalid date: %v", err)
	}

	amount, err := strconv.ParseFloat(record[3], 64)
	if err != nil {
		return model.Expense{}, fmt.Errorf("invalid amount: %v", err)
	}

	return model.Expense{
		ID:          id,
		Date:        date,
		Description: record[2],
		Amount:      amount,
	}, nil
}

func ensureDir(dirName string) error {
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		if err := os.MkdirAll(dirName, 0755); err != nil {
			return fmt.Errorf("error creating directory: %v", err)
		}
	}
	return nil
}

func SaveExpenses(expenses []model.Expense) error {
	if err := ensureDir("results"); err != nil {
		fmt.Println("Error creating directory:", err)
		return err
	}

	file, err := os.Create(expanseFile)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}
	defer file.Close()

	return writeCSV(file, expenses)
}

func writeCSV(file *os.File, expenses []model.Expense) error {
	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, exp := range expenses {
		record := expenseToCSVRecord(exp)
		if err := writer.Write(record); err != nil {
			fmt.Println("Error writing record:", err)
			return err
		}
	}
	return nil
}

func expenseToCSVRecord(exp model.Expense) []string {
	return []string{
		strconv.Itoa(exp.ID),
		exp.Date.Format("2006-01-02"),
		exp.Description,
		fmt.Sprintf("%.2f", exp.Amount),
	}
}

func CalculateTotal(expenses []model.Expense, month int) float64 {
	var total float64
	for _, expense := range expenses {
		if isInMonth(expense.Date, month) {
			total += expense.Amount
		}
	}
	return total
}

func isInMonth(date time.Time, month int) bool {
	return month <= 0 || date.Month() == time.Month(month)
}
