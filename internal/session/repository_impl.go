package session

import (
	"context"

	sharederrors "github.com/juanchrstian/restaurant-api/internal/shared/errors"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(
	db *gorm.DB,
) Repository {

	return &repository{
		db: db,
	}

}

func (r *repository) Create(
	ctx context.Context,
	session *Session,
) error {

	return r.db.
		WithContext(ctx).
		Create(session).
		Error

}

func (r *repository) GetActive(
	ctx context.Context,
) (*Session, error) {

	var session Session

	err := r.db.
		WithContext(ctx).
		Where(
			"status = ?",
			StatusOpen,
		).
		First(&session).
		Error

	if err != nil {

		if err == gorm.ErrRecordNotFound {
			return nil, sharederrors.ErrSessionNotFound
		}

		return nil, err
	}

	return &session, nil

}

func (r *repository) Update(
	ctx context.Context,
	session *Session,
) error {

	return r.db.
		WithContext(ctx).
		Save(session).
		Error

}

func (r *repository) GetSummary(
	ctx context.Context,
	sessionID string,
) (*Summary, error) {

	var summary Summary

	err := r.db.
		WithContext(ctx).
		Table("orders").
		Select(`
			COUNT(*) AS total_orders,

			COALESCE(SUM(total_amount),0) AS gross_sales,

			COALESCE(SUM(
				CASE
					WHEN payment_method = 'CASH'
					THEN total_amount
					ELSE 0
				END
			),0) AS cash_sales,

			COALESCE(SUM(
				CASE
					WHEN payment_method = 'QRIS'
					THEN total_amount
					ELSE 0
				END
			),0) AS qris_sales,

			COALESCE(SUM(
				CASE
					WHEN payment_method = 'DEBIT'
					THEN total_amount
					ELSE 0
				END
			),0) AS debit_sales
		`).
		Where(
			"session_id = ? AND status = ?",
			sessionID,
			"PAID",
		).
		Scan(&summary).Error

	if err != nil {
		return nil, err
	}

	return &summary, nil
}

func (r *repository) GetAll(
	ctx context.Context,
) ([]Session, error) {

	var sessions []Session

	err := r.db.
		WithContext(ctx).
		Order("opened_at DESC").
		Find(&sessions).Error

	if err != nil {
		return nil, err
	}

	return sessions, nil
}
