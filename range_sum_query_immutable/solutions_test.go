package range_sum_query_immutable

import "testing"

func TestNumArray_SumRange(t *testing.T) {
	type args struct {
		left  int
		right int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		initial []int
	}{
		{
			name: "test1",
			args: args{
				left:  0,
				right: 2,
			},
			initial: []int{-2, 0, 3, -5, 2, -1},
			want:    1,
		},
		{
			name: "test2",
			args: args{
				left:  2,
				right: 5,
			},
			initial: []int{-2, 0, 3, -5, 2, -1},
			want:    -1,
		},
		{
			name: "test3",
			args: args{
				left:  0,
				right: 5,
			},
			initial: []int{-2, 0, 3, -5, 2, -1},
			want:    -3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := Constructor(tt.initial)
			if got := this.SumRange(tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("SumRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
