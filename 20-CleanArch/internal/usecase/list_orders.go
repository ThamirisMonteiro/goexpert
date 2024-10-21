package usecase

import "github.com/devfullcycle/20-CleanArch/internal/entity"

type ListOrdersInputDTO struct{}

type ListOrdersOutputDTO struct {
	Orders []OrderOutputDTO `json:"orders"`
}

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(orderRepository entity.OrderRepositoryInterface) *ListOrdersUseCase {
	return &ListOrdersUseCase{OrderRepository: orderRepository}
}

func (c *ListOrdersUseCase) Execute(_ ListOrdersInputDTO) (ListOrdersOutputDTO, error) {
	orders, err := c.OrderRepository.List()
	if err != nil {
		return ListOrdersOutputDTO{}, err
	}

	var ordersDTO []OrderOutputDTO
	for _, order := range orders {
		ordersDTO = append(ordersDTO, OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}

	return ListOrdersOutputDTO{Orders: ordersDTO}, nil
}
