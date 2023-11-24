package IncomeAndExpense

type IncomeAndExpensesList struct {
	TotalBalance          float32                     `json:"totalBalance"`
	IncomeAndExpensesList []IncomeAndExpensesPerMonth `json:"incomeAndExpensesList"`
}

type IncomeAndExpensesPerMonth struct {
	Month                 string              `json:"month"`
	IncomeAndExpensesList []IncomeAndExpenses `json:"incomeAndExpenses"`
}

type IncomeAndExpenses struct {
	ID          int64    `json:"ID"`
	Day         string   `json:"day"`
	Time        string   `json:"time"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Amount      float32  `json:"amount"`
	Type        string   `json:"type"`
	Category    Category `json:"category"`
}

type Category struct {
	ID   int64  `json:"ID"`
	Name string `json:"name"`
}
