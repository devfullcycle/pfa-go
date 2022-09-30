package usecase

import (
	"github.com/devfullcycle/pfa-go/internal/order/entity"
)

type GetTotalOutputDto struct {
	Total int
}

type GetTotalUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewGetTotalUseCase(orderRepository entity.OrderRepositoryInterface) *GetTotalUseCase {
	return &GetTotalUseCase{OrderRepository: orderRepository}
}

func (c *GetTotalUseCase) Execute() (*GetTotalOutputDto, error) {
	total, err := c.OrderRepository.GetTotal()
	if err != nil {
		return nil, err
	}
	return &GetTotalOutputDto{Total: total}, nil
}
