package dotpattern

import (
	"testing"
)

func TestReorderBits(t *testing.T) {
	type args struct {
		src   byte
		order [8]uint8
	}
	tests := []struct {
		name    string
		args    args
		wantDst byte
	}{
		{
			name:    "basic",
			args:    args{1, [8]uint8{7, 6, 5, 4, 3, 2, 1, 0}},
			wantDst: 128,
		},
		{
			name:    "reverse",
			args:    args{240, [8]uint8{7, 6, 5, 4, 3, 2, 1, 0}},
			wantDst: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDst := ReorderBits(tt.args.src, tt.args.order); gotDst != tt.wantDst {
				t.Errorf("ReorderBits() = %v, want %v", gotDst, tt.wantDst)
			}
		})
	}
}

func TestDefaultOrder(t *testing.T) {
	test := []byte{1, 2, 4, 8, 16, 32, 64, 128, 15, 240, 102, 6, 96}
	result := Bytes(test, DefaultOrder)
	expect := "⠁⠂⠄⡀⠈⠐⠠⢀⡇⢸⠶⠆⠰"
	if result != expect {
		t.Errorf("Bytes() =%v, want %v", result, expect)
	}
}
