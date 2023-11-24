package dataModel

type DetailIncomeAndExpense struct {
	ID          int64   `gorm:"column:ID"`
	Title       string  `gorm:"column:title"`
	Description string  `gorm:"column:description"`
	Amount      float32 `gorm:"column:amount"`
	CategoryID  int64   `gorm:"column:categoryID"`
	PrentID     int64   `gorm:"column:parentID"`
}
