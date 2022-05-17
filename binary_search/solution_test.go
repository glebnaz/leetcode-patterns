package binary_search

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 9,
			},
			want: 4,
		},
		{
			name: "test2",
			args: args{
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 2,
			},
			want: -1,
		},
		{
			name: "test3",
			args: args{
				nums:   []int{5},
				target: 5,
			},
			want: 0,
		},
		{
			name: "test4",
			args: args{
				nums:   []int{2, 5},
				target: 2,
			},
			want: 0,
		},
		{
			name: "test5",
			args: args{
				nums:   []int{-1, 0, 5},
				target: 5,
			},
			want: 2,
		},
		{
			name: "test6",
			args: args{
				nums:   []int{1, 2, 3, 4, 5, 6, 7, 8},
				target: 5,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
