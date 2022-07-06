package service

import (
	"context"
	"encoding/json"
	"enigmacamp.com/lopei_grpc_srv/repository"
)

type LopeiService struct {
	repo repository.LopeiRepository
}

func (s *LopeiService) mustEmbedUnimplementedLopeiPaymentServer() {
	panic("implement me")
}

func (s *LopeiService) CheckBalance(ctx context.Context, in *CheckBalanceMessage) (*ResultMessage, error) {
	lopeiId := in.LopeiId
	customer, err := s.repo.RetrieveById(lopeiId)
	if err != nil {
		return nil, err
	}
	c, err := json.Marshal(customer)
	if err != nil {
		return nil, err
	}
	resultMessage := &ResultMessage{
		Result: string(c),
		Error:  nil,
	}

	return resultMessage, nil
}
func (s *LopeiService) DoPayment(ctx context.Context, in *PaymentMessage) (*ResultMessage, error) {
	lopeiId := in.LopeiId
	requestDeduction := in.Amount
	customer, err := s.repo.RetrieveById(lopeiId)
	if err != nil {
		return nil, err
	}
	if customer.Balance < requestDeduction {
		return &ResultMessage{
			Result: "FAILED",
			Error: &Error{
				Code:    "X07",
				Message: "Insufficient Balance",
			},
		}, nil
	}

	resultMessage := &ResultMessage{
		Result: "SUCCESS",
		Error:  nil,
	}

	return resultMessage, nil
}

func NewLopeiService(repo repository.LopeiRepository) *LopeiService {
	service := new(LopeiService)
	service.repo = repo
	return service
}
