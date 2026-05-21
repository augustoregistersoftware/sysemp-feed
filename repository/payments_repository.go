package repository

import (
	"context"
	"sysemp_travel/model"
)

type PaymentsRepository struct {
	*Repository
}

func NewPaymentsRepository(baseRepo *Repository) PaymentsRepository {
	return PaymentsRepository{
		Repository: baseRepo,
	}
}

func (r *PaymentsRepository) GetPayments(ctx context.Context) ([]model.Payment, error) {
	var payment []model.Payment
	rows, err := r.DB.QueryContext(ctx, "SELECT id_payments, name FROM payments")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p model.Payment
		if err := rows.Scan(&p.ID, &p.Name); err != nil {
			return nil, err
		}
		payment = append(payment, p)
	}

	return payment, nil
}
