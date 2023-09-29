package binary_search

import "testing"

func Test_binarySearch(t *testing.T) {
	type args struct {
		sortedArray []int
		target      int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{
		{
			"Empty array, search 1, (-1, false) is expected",
			args{
				[]int{},
				1,
			},
			-1,
			false,
		},
		{
			"Empty array, search 1, (-1, false) is expected",
			args{
				nil,
				1,
			},
			-1,
			false,
		},
		{
			"array [1,4,6,10,13,14,20,100,120,200,1000,1500,100000], search 10, (3, true) is expected",
			args{
				[]int{1, 4, 6, 10, 13, 14, 20, 100, 120, 200, 1000, 1500, 100000},
				10,
			},
			3,
			true,
		},
		{
			"array [1,4,6,10,13,14,20,100,120,200,1000,1500,100000], search 1002, (-1, false) is expected",
			args{
				[]int{1, 4, 6, 10, 13, 14, 20, 100, 120, 200, 1000, 1500, 100000},
				1002,
			},
			-1,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := binarySearch(tt.args.sortedArray, tt.args.target)
			if got != tt.want {
				t.Errorf("binarySearch() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("binarySearch() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
