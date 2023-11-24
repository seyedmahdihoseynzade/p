package dataModel

type Category struct {
	ID       int64  `gorm:"column:ID"`
	UserID   int64  `gorm:"column:userID"`
	Name     string `gorm:"column:name"`
	ParentID int64  `gorm:"column:parentID"`
}
