package session

type Summary struct {
	TotalOrders int `gorm:"column:total_orders"`

	GrossSales int64 `gorm:"column:gross_sales"`

	CashSales int64 `gorm:"column:cash_sales"`

	QRISSales int64 `gorm:"column:qris_sales"`

	DebitSales int64 `gorm:"column:debit_sales"`
}
