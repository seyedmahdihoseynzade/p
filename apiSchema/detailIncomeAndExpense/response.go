package detailIncomeAndExpense

type IncomeAndExpensesList struct {
	Budget                float32             `json:"budget"`
	Month                 string              `json:"month"`
	Name                  string              `json:"name"`
	IncomeAndExpensesList []IncomeAndExpenses `json:"detailIncomeAndExpensesList"`
}

type IncomeAndExpenses struct {
	Day         string  `json:"day"`
	Time        string  `json:"time"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Amount      float32 `json:"amount"`
	Type        string  `json:"type"`
	Category    string  `json:"category"`
}
