package service

import (
	"go_demo/model"
	"testing"
)

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
/*
go test -v -run='none' -benchtime='3s' -bench='BenchmarkDemoZ'
*/
func BenchmarkDemoZ(b *testing.B) {
	var rec []model.Order
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t1("test_name", &rec, 1, 30)
		rec[0].Status = "111"
	}
}

func BenchmarkDemoZ2(b *testing.B) {
	var rec []*model.Order
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t2("test_name", &rec, 1, 30)
		rec[0].Status = "111"
	}
}
/*
goos: linux
goarch: amd64
pkg: go_demo/service
BenchmarkDemoZ-4   	    2000	    541159 ns/op
PASS

goos: linux
goarch: amd64
pkg: go_demo/service
BenchmarkDemoZ2-4   	    2000	    546862 ns/op
PASS
*/

func t1(username string, rec *[]model.Order, pageNo uint, size uint) (recordCnt uint, pageCnt uint, err error) {
	if pageNo == 0 {
		pageNo = 1
	}
	dbConn.Model(&model.Order{}).Count(&recordCnt)
	err = dbConn.Where("user_name = ?", username).Limit(size).Offset((pageNo - 1) * size).Find(rec).Error
	pageCnt = (recordCnt + size - 1) / size
	return
}

func t2(username string, rec *[]*model.Order, pageNo uint, size uint) (recordCnt uint, pageCnt uint, err error) {
	if pageNo == 0 {
		pageNo = 1
	}
	dbConn.Model(&model.Order{}).Count(&recordCnt)
	err = dbConn.Where("user_name = ?", username).Limit(size).Offset((pageNo - 1) * size).Find(rec).Error
	pageCnt = (recordCnt + size - 1) / size
	return
}