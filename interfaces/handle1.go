package interfaces

import (
	"context"

	"github.com/82wutao/ee-rpcdeclare/order"
	"github.com/82wutao/ee-services/services"
)

type OrderService int

func (os *OrderService) HandleName() string {
	return "order-service"
}

func (os *OrderService) OrderQuery(ctx context.Context, req *order.OrderQueryReq, resp *order.OrderQueryResp) error {
	cmd := services.NewCommand([]interface{}{req.UserID}, func(args []interface{}) ([]interface{}, error) {
		userID := args[0].(int)
		orders := services.OrderQuery(userID)
		return []interface{}{orders}, nil
	})
	services.SyncOneCommand(cmd)
	if cmd.Error != nil {
		return cmd.Error
	}
	orders := cmd.Return[0].([]int)
	resp.Orders = orders
	return nil
}

func (os *OrderService) OrderSubmit(ctx context.Context, req *order.OrderSubmitReq, resp *order.OrderSubmitResp) error {
	cmd := services.NewCommand([]interface{}{req.UserID}, func(args []interface{}) ([]interface{}, error) {
		userID := args[0].(int)
		newOrderID := services.OrderSubmit(userID, nil)
		return []interface{}{newOrderID}, nil
	})
	services.SyncOneCommand(cmd)
	if cmd.Error != nil {
		return cmd.Error
	}
	order := cmd.Return[0].(int)
	resp.OrderID = order
	resp.UserID = req.UserID
	return nil
}
func (os *OrderService) OrderCancel(ctx context.Context, req *order.OrderCancelReq, resp *order.OrderCancelResp) error {
	cmd := services.NewCommand([]interface{}{req.UserID}, func(args []interface{}) ([]interface{}, error) {
		userID := args[0].(int)
		orderID := args[1].(int)
		suc := services.OrderCancel(userID, orderID)
		return []interface{}{suc}, nil
	})
	services.SyncOneCommand(cmd)
	if cmd.Error != nil {
		return cmd.Error
	}
	suc := cmd.Return[0].(bool)
	resp.Suc = suc
	return nil
}
