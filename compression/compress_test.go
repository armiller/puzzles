package compress

import "testing"

func Test_compress(t *testing.T) {
	type args struct {
		uncompressed string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			"Empty string",
			args{uncompressed: ""},
			"",
		},
		{
			"one character",
			args{uncompressed: "a"},
			"a1",
		},
		{
			"two characters",
			args{uncompressed: "aa"},
			"a2",
		},
		{
			"two unique characters",
			args{uncompressed: "ab"},
			"a1b1",
		},
		{
			"sequential",
			args{uncompressed: "aaaabbbbccccd"},
			"a4b4c4d1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compress(tt.args.uncompressed); got != tt.want {
				t.Errorf("compress() = %v, want %v", got, tt.want)
			}
		})
	}
}
