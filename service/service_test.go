package service

import "testing"

func TestAddNewOrder(t *testing.T) {
	type args struct {
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
		{"test1", args{"t1", 12.52, "PASS", ""}, false},
		{"test2", args{"t2", 12.45, "FAIL", ""}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AddNewOrder(tt.args.username, tt.args.amount, tt.args.status, tt.args.fileURL); (err != nil) != tt.wantErr {
				t.Errorf("AddNewOrder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
