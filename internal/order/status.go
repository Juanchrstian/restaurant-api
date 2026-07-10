package order

type OrderStatus string

const (
	OrderPending   OrderStatus = "PENDING"
	OrderPaid      OrderStatus = "PAID"
	OrderCancelled OrderStatus = "CANCELLED"
)

type PaymentMethod string

const (
	PaymentCash  PaymentMethod = "CASH"
	PaymentQRIS  PaymentMethod = "QRIS"
	PaymentDebit PaymentMethod = "DEBIT"
)
