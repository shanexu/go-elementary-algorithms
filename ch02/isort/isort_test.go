package isort

import (
	"reflect"
	"testing"
)

func TestISort(t *testing.T) {
	type args struct {
		xs []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"nil",
			args{nil},
			nil,
		},
		{
			"single",
			args{[]int{1}},
			[]int{1},
		},
		{
			"two",
			args{[]int{2, 1}},
			[]int{1, 2},
		},
		{
			"three",
			args{[]int{2, 3, 1}},
			[]int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if ISort(tt.args.xs); !reflect.DeepEqual(tt.args.xs, tt.want) {
				t.Errorf("ISort() = %v, want %v", tt.args.xs, tt.want)
			}
		})
	}
}
