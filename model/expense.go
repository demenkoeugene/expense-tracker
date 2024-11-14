package model

import (
	"fmt"
	"time"
)

type Expense struct {
	ID          int
	Date        time.Time
	Description string
	Amount      float64
}

func (e Expense) Print() {
	fmt.Printf("# %-4d %-12s %-20s $%6.2f\n",
		e.ID, e.Date.Format("2006-01-02"), e.Description, e.Amount)
}
