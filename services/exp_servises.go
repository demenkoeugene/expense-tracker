package services

import (
	"encoding/csv"
	"expense-tracker/model"
	"fmt"
	"os"
	"strconv"
	"time"
)

const expanseFile = "results/tasks.csv"

func LoadExpanseList() ([]model.Expense, error) {
	file, err := os.Open(expanseFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []model.Expense{}, nil
		}
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var expenses []model.Expense
	for _, record := range records {
		id, _ := strconv.Atoi(record[0])
		date, _ := time.Parse("2006-01-02", record[1])
		amount, _ := strconv.ParseFloat(record[3], 64)
		expenses = append(expenses, model.Expense{
			ID:          id,
			Date:        date,
			Description: record[2],
			Amount:      amount,
		})
	}
	return expenses, nil
}

func ensureDir(dirName string) error {
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		err := os.MkdirAll(dirName, 0755)
		if err != nil {
			return fmt.Errorf("Error creating directory: %v", err)
		}
	}
	return nil
}

func SaveExpenses(expenses []model.Expense) error {
	err := ensureDir("results")
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return err
	}

	file, err := os.Create(expanseFile)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	for _, exp := range expenses {
		record := []string{
			strconv.Itoa(exp.ID),
			exp.Date.Format("2006-01-02"),
			exp.Description,
			fmt.Sprintf("%.2f", exp.Amount),
		}
		writer.Write(record)
	}
	return nil
}

func CalculateTotal(expenses []model.Expense, month int) float64 {
	var total float64
	for _, expense := range expenses {
		if month > 0 && expense.Date.Month() != time.Month(month) {
			continue
		}
		total += expense.Amount
	}
	return total
}
