package repository

import "enigmacamp.com/lopei_grpc_srv/model"

type LopeiRepository interface {
	RetrieveById(id int32) (model.Customer, error)
}

type lopeiRepository struct {
	db []model.Customer
}

func (p *lopeiRepository) RetrieveById(id int32) (model.Customer, error) {
	for _, cust := range p.db {
		if cust.LopeiId == id {
			return cust, nil
		}
	}
	return model.Customer{}, nil
}

func NewLopeiRepository() LopeiRepository {
	repo := new(lopeiRepository)
	repo.db = []model.Customer{
		{LopeiId: 1, Balance: 1000},
		{LopeiId: 2, Balance: 3000},
		{LopeiId: 3, Balance: 500},
	}
	return repo
}
