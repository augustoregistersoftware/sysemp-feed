package usecase

import (
	"context"
	"sysemp_travel/model"
	"sysemp_travel/repository"
)

type PaymentsUseCase struct {
	repository repository.PaymentsRepository
}

func NewPaymentsUseCase(paymentsRepo repository.PaymentsRepository) PaymentsUseCase {
	return PaymentsUseCase{
		repository: paymentsRepo,
	}
}

func (u *PaymentsUseCase) GetPayments(ctx context.Context) ([]model.Payment, error) {
	payment, err := u.repository.GetPayments(ctx)
	if err != nil {
		return nil, err
	}
	return payment, nil
}
