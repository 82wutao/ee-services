package interfaces_test

import (
	"context"
	"testing"

	"github.com/82wutao/ee-rpcdeclare/order"
	"github.com/82wutao/ee-services/interfaces"
	"github.com/82wutao/ee-services/services"
)

func TestOrderService_Submit(t *testing.T) {
	services.Loop(50)

	os := new(interfaces.OrderService)

	type args struct {
		ctx  context.Context
		req  *order.OrderSubmitReq
		resp *order.OrderSubmitResp
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "submit1",
			args: args{
				ctx:  context.Background(),
				req:  &order.OrderSubmitReq{UserID: 2, OrderParam: nil},
				resp: &order.OrderSubmitResp{UserID: 2, OrderID: 201},
			},
			wantErr: false,
		},
		{
			name: "submit2",
			args: args{
				ctx:  context.Background(),
				req:  &order.OrderSubmitReq{UserID: 2, OrderParam: nil},
				resp: &order.OrderSubmitResp{UserID: 2, OrderID: 202},
			},
			wantErr: false,
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := os.Submit(tt.args.ctx, tt.args.req, tt.args.resp); (err != nil) != tt.wantErr {
				t.Errorf("OrderService.Submit() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.args.resp.OrderID != tt.args.req.UserID*100+i+1 {
				t.Errorf("OrderService.Submit() orderID != UserID*100+len(orders)+1, %d", tt.args.resp.OrderID)
			}
		})
	}
	services.StopLoop()
}
