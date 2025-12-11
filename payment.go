package domain

type Payment struct {
	ID      int64 `gorm:"primaryKey"`
	OrderID int64
	UserID  int64
	Total   float32
	BillID  int64
}
