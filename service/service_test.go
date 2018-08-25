package service

import (
	"go_demo/model"
	"testing"
)

var orderid1 = "1535197569163105712461074c783ac4441a8cd2efa63d39dde"
var orderid2 = "15351975691743669344e0e284091b940148a0fbc8fb3d62a5c"
var username = "test_name"

func TestAddNewOrder(t *testing.T) {
	type args struct {
		orderid  string
		username string
		amount   float64
		status   string
		fileURL  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"testpass", args{orderid1, username, 12.33, "PASS", "xx"}, false},
		{"testpass", args{orderid2, username, 11.22, "FAIL", "yy"}, false},
		{"testfail", args{"", "", 0.0, "", ""}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AddNewOrder(tt.args.orderid, tt.args.username, tt.args.amount, tt.args.status, tt.args.fileURL); (err != nil) != tt.wantErr {
				t.Errorf("AddNewOrder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdateOrder(t *testing.T) {
	type args struct {
		orderid string
		amount  float64
		status  string
		fileURL string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"testpass", args{orderid1, 99.99, "FAIL", "xx"}, false},
		{"testfail", args{orderid1, 0.0, "PASS", ""}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateOrder(tt.args.orderid, tt.args.amount, tt.args.status, tt.args.fileURL); (err != nil) != tt.wantErr {
				t.Errorf("UpdateOrder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdateOrderFileURL(t *testing.T) {
	type args struct {
		orderid string
		fileURL string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"testpass", args{orderid1, "yy"}, false},
		{"testfail", args{orderid1, ""}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateOrderFileURL(tt.args.orderid, tt.args.fileURL); (err != nil) != tt.wantErr {
				t.Errorf("UpdateOrderFileURL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetOrder(t *testing.T) {
	type args struct {
		rec *model.Order
	}
	rec := model.Order{OrderId: orderid1}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"t1", args{&rec}, false},
		{"t2", args{&rec}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetOrder(tt.args.rec); (err != nil) != tt.wantErr {
				t.Errorf("GetOrder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetOrderListByUserName(t *testing.T) {
	type args struct {
		key string
		rec []model.Order
	}

	var rec []model.Order

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"t1", args{"test", rec}, false},
		{"t2", args{"xx", rec}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetOrderListByUserName(tt.args.key, tt.args.rec); (err != nil) != tt.wantErr {
				t.Errorf("GetOrderListByUserName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetOrderFileURL(t *testing.T) {
	type args struct {
		orderid string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"testpass", args{orderid2}, "yy", false},
		{"testpass", args{""}, "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetOrderFileURL(tt.args.orderid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOrderFileURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetOrderFileURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
