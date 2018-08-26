package service

import "testing"

func Test_transAmount(t *testing.T) {
	type args struct {
		from   string
		to     string
		amount float64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"testpass", args{orderid1, orderid2, 2}, false},
		{"testfail", args{orderid1, orderid2, 45}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := transAmount(tt.args.from, tt.args.to, tt.args.amount); (err != nil) != tt.wantErr {
				t.Errorf("transAmount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
