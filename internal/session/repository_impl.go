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
