package models

import (
	"reflect"
	"testing"
)

func TestCreateUseDetail(t *testing.T) {
	targetLine := `"****-****-****-1234　ｘｘｘ　ｘｘ　ｘｘ様","≪ショッピング取組（国内）≫","2022/12/01","決済Ａ","2,653","１回","","","275","国内","","*"`
	type args struct {
		source string
		month  string
		line   string
	}
	tests := []struct {
		name string
		args args
		want UseDetail
	}{
		{"normal",
			args{"TEST", "202501", targetLine},
			UseDetail{Source: "TEST", PayMonth: "202501", Date: "2022/12/01", Title: "決済Ａ", Amount: 2653, Line: targetLine},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateUseDetail(tt.args.source, tt.args.month, tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateUseDetail() = %v, want %v", got, tt.want)
			}
		})
	}
}
