package interfaces

import "context"

type OrderService int

func (os *OrderService) HandleName() string {
	return "order-service"
}

type Req struct {
	I int
}
type Resp struct {
	O int
}

func (os *OrderService) Query(ctx context.Context, req *Req, resp *Resp) error {
	return nil
}
func (os *OrderService) Submit(ctx context.Context, req *Req, resp *Resp) error {
	return nil
}
func (os *OrderService) Cancel(ctx context.Context, req *Req, resp *Resp) error {
	return nil
}
