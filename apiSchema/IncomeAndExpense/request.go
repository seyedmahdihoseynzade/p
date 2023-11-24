package IncomeAndExpense

type Create struct {
	Title       string  `json:"title"`
	Amount      float32 `json:"amount"`
	Description string  `json:"description"`
	Category    int64   `json:"categoryID"`
}
