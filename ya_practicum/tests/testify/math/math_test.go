// Generated TestAdd
package math

import "testing"

func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// Put your testcases here
		{
			name: "Test Positive",
			args: args{
				a: 1,
				b: 2,
			},
			want:    3,
			wantErr: false,
		},
		{
			name: "Test Negative 1",
			args: args{
				a: -1,
				b: 2,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Test Negative 2",
			args: args{
				a: 1,
				b: -2,
			},
			want:    0,
			wantErr: true,
		},

		{
			name: "Test Negative all",
			args: args{
				a: -1,
				b: -2,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Test Zero 1",
			args: args{
				a: 0,
				b: 1,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Add(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
