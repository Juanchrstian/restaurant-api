package session

import "time"

type OpenSessionRequest struct {
	OpenedBy    string `json:"opened_by" validate:"required,max=100"`
	OpeningCash int64  `json:"opening_cash" validate:"gte=0"`
}

type CloseSessionRequest struct {
	ClosingCash int64 `json:"closing_cash" validate:"gte=0"`
}

type SessionResponse struct {
	ID          string  `json:"id"`
	Status      string  `json:"status"`
	OpenedBy    string  `json:"opened_by"`
	OpeningCash int64   `json:"opening_cash"`
	ClosingCash *int64  `json:"closing_cash,omitempty"`
	OpenedAt    string  `json:"opened_at"`
	ClosedAt    *string `json:"closed_at,omitempty"`
}

func ToResponse(
	session *Session,
) SessionResponse {

	response := SessionResponse{

		ID: session.ID.String(),

		Status: string(session.Status),

		OpenedBy: session.OpenedBy,

		OpeningCash: session.OpeningCash,

		ClosingCash: session.ClosingCash,

		OpenedAt: session.OpenedAt.Format(time.RFC3339),
	}

	if session.ClosedAt != nil {

		value := session.ClosedAt.Format(time.RFC3339)

		response.ClosedAt = &value

	}

	return response

}

type CloseSessionResponse struct {
	SessionID string `json:"session_id"`

	OpenedBy string `json:"opened_by"`

	Status string `json:"status"`

	OpenedAt time.Time `json:"opened_at"`

	ClosedAt *time.Time `json:"closed_at"`

	OpeningCash int64 `json:"opening_cash"`

	ClosingCash int64 `json:"closing_cash"`

	TotalOrders int `json:"total_orders"`

	GrossSales int64 `json:"gross_sales"`

	CashSales int64 `json:"cash_sales"`

	QRISSales int64 `json:"qris_sales"`

	DebitSales int64 `json:"debit_sales"`
}

func ToCloseSessionResponse(
	session *Session,
	summary *Summary,
) *CloseSessionResponse {

	var closingCash int64

	if session.ClosingCash != nil {
		closingCash = *session.ClosingCash
	}

	return &CloseSessionResponse{
		SessionID: session.ID.String(),

		OpenedBy: session.OpenedBy,

		Status: string(session.Status),

		OpenedAt: session.OpenedAt,

		ClosedAt: session.ClosedAt,

		OpeningCash: session.OpeningCash,

		ClosingCash: closingCash,

		TotalOrders: summary.TotalOrders,

		GrossSales: summary.GrossSales,

		CashSales: summary.CashSales,

		QRISSales: summary.QRISSales,

		DebitSales: summary.DebitSales,
	}
}
