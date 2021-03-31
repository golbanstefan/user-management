package response

import (
	"reflect"
	"testing"
)

func TestGet(t *testing.T) {
	type args struct {
		d interface{}
	}
	tests := []struct {
		name string
		args args
		want Response
	}{
		{"1", args{
			"Test",
		}, ResponseStruct{
			Data: "Test",
		}},
		{"2", args{
			"Test 2",
		}, ResponseStruct{
			Data: "Test 2",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Get(tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
