package path_sum

import "testing"

func Test_hasPathSum(t *testing.T) {
	type args struct {
		root      *TreeNode
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 11,
							Left: &TreeNode{
								Val: 7,
							},
						},
					},
				},
				targetSum: 27,
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 11,
							Left: &TreeNode{
								Val: 7,
							},
							Right: &TreeNode{
								Val:   2,
								Right: nil,
								Left:  nil,
							},
						},
					},
					Right: &TreeNode{
						Val: 8,
						Left: &TreeNode{
							Val:   13,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val: 4,
							Right: &TreeNode{
								Val:   1,
								Left:  nil,
								Right: nil,
							},
						},
					},
				},
				targetSum: 22,
			},
			want: true,
		},
		{
			name: "test3",
			args: args{
				root: &TreeNode{
					Val:   1,
					Right: nil,
					Left:  nil,
				},
				targetSum: 1,
			},
			want: true,
		},
		{
			name: "test4",
			args: args{
				root: &TreeNode{
					Val:   1,
					Right: &TreeNode{2, nil, nil},
					Left:  nil,
				},
				targetSum: 1,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPathSum(tt.args.root, tt.args.targetSum); got != tt.want {
				t.Errorf("hasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
