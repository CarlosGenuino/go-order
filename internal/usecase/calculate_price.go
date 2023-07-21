package usecase

import "github.com/carlosgenuino/fieldservice/internal/entity"

type OrderInput struct {
	Id    string  `json:"id"`
	Price float64 `json:"price"`
	Tax   float64 `json:"tax"`
}

type OrderOutput struct {
	Id         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type CalculateFinalPrice struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewUseCaseCalculateFinalPrice(repository entity.OrderRepositoryInterface) *CalculateFinalPrice {
	return &CalculateFinalPrice{
		OrderRepository: repository,
	}
}

func (c *CalculateFinalPrice) Execute(input OrderInput) (*OrderOutput, error) {
	order, err := entity.NewOrder(input.Id, input.Price, input.Tax)
	if err != nil {
		return nil, err
	}
	err = order.CalculateFinalPrice()
	if err != nil {
		return nil, err
	}
	err = c.OrderRepository.Save(order)
	if err != nil {
		return nil, err
	}
	return &OrderOutput{
		Id:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}, nil
}
