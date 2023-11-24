package detailIncomeAndExpense

type Create struct {
	Title                    string  `json:"title"`
	Amount                   float32 `json:"amount"`
	Description              string  `json:"description"`
	Category                 int64   `json:"categoryID"`
	ParentIncomeAndExpenseID int64   `json:"parentIncomeAndExpenseID"`
}

type List struct {
	ParentID int64 `json:"parentID"`
}
