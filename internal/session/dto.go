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
