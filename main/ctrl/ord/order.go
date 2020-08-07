package ord

type IOrderService interface {
	findOrder() string
	saveOrder() string
}

type OrderService struct {
}

func (o *OrderService) findOrder() string {
	return o.T
}

func (o *OrderService) saveOrder() string {
	o.Status = "Saved"
	return "saved"
}
