package estimate

import "testing"

func TestEstimateValue(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "Test value < 10",
			args: args{
				value: 2,
			},
			want: "small",
		},
		{
			name: "Test value > 10 and < 100",
			args: args{
				value: 99,
			},
			want: "medium",
		},
		{
			name: "Test value > 100",
			args: args{
				value: 101,
			},
			want: "big",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EstimateValue(tt.args.value); got != tt.want {
				t.Errorf("EstimateValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
