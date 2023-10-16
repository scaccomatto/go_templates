package grpc

import (
	"paru.net/exagonalapp/internal/ports"
)

type Adapter struct {
	api  ports.APIPort
	port int
	//order.UnimplementedOrderServer
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{api: api, port: port}
}

// func (a Adapter) Create(ctx context.Context, request *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
// 	var orderItems []domain.OrderItem
// 	for _, orderItem := range request.OrderItems {
// 		orderItems = append(orderItems, domain.OrderItem{
// 			ProductCode: orderItem.ProductCode,
// 			UnitPrice:   orderItem.UnitPrice,
// 			Quantity:    orderItem.Quantity,
// 		})
// 	}
// 	newOrder := domain.NewOrder(request.UserId, orderItems)
// 	result, err := a.api.PlaceOrder(newOrder)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &order.CreateOrderResponse{OrderId: result.ID}, nil
// }
