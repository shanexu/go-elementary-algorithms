package isort

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
			ISortB(tt.args.xs)
			assert.Equal(t, tt.want, tt.args.xs)
		})
	}
}

func TestBinarySearch(t *testing.T) {
	type args struct {
		xs []int
		x  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"nil",
			args{nil, 1},
			0,
		},
		{
			"single",
			args{[]int{1}, 1},
			0,
		},
		{
			"single2",
			args{[]int{1}, 0},
			0,
		},
		{
			"single3",
			args{[]int{1}, 2},
			1,
		},
		{
			"two",
			args{[]int{1, 2}, 2},
			1,
		},
		{
			"two1",
			args{[]int{1, 2}, 3},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.args.xs, tt.args.x); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestISortB(t *testing.T) {
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
			ISortB(tt.args.xs)
			assert.Equal(t, tt.want, tt.args.xs)
		})
	}
}
