package dataModel

type IncomeAndExpense struct {
	ID          int64   `gorm:"column:ID"`
	UserID      int64   `gorm:"column:userID"`
	Date        int64   `gorm:"column:date"`
	Title       string  `gorm:"column:title"`
	Description string  `gorm:"column:description"`
	Amount      float32 `gorm:"column:amount"`
	CategoryID  int64   `gorm:"column:categoryID"`
}
