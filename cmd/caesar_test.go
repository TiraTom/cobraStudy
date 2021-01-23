package cmd

import (
	"reflect"
	"testing"
)

func Test_caesar(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "大文字が含まれる場合", args: args{args: []string{"DEF"}}, want: []string{"cde", "bcd", "abc", "zab", "yza", "xyz", "wxy", "vwx", "uvw", "tuv", "stu", "rst", "qrs", "pqr", "opq", "nop", "mno", "lmn", "klm", "jkl", "ijk", "hij", "ghi", "fgh", "efg"}},
		{name: "小文字だけの場合", args: args{args: []string{"def"}}, want: []string{"cde", "bcd", "abc", "zab", "yza", "xyz", "wxy", "vwx", "uvw", "tuv", "stu", "rst", "qrs", "pqr", "opq", "nop", "mno", "lmn", "klm", "jkl", "ijk", "hij", "ghi", "fgh", "efg"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := caesar(tt.args.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("caesar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_shiftASCIICODEs(t *testing.T) {
	type args struct {
		runes []rune
		shift int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "単純にずらせば良いパターン1", args: args{runes: []rune("def"), shift: 3}, want: "abc"},
		{name: "zからの計算がいるパターン1", args: args{runes: []rune("zae"), shift: 3}, want: "wxb"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shiftASCIICODEs(tt.args.runes, tt.args.shift); got != tt.want {
				t.Errorf("shiftASCIICODEs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_shiftASCIICODE(t *testing.T) {
	type args struct {
		base  rune
		shift int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "単純にずらせば良いパターン1", args: args{base: []rune("e")[0], shift: 3}, want: int([]rune("b")[0])},
		{name: "単純にずらせば良いパターン2", args: args{base: []rune("z")[0], shift: 2}, want: int([]rune("x")[0])},
		{name: "zからの計算があるパターン1", args: args{base: []rune("b")[0], shift: 2}, want: int([]rune("z")[0])},
		{name: "zからの計算があるパターン2", args: args{base: []rune("a")[0], shift: 1}, want: int([]rune("z")[0])},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shiftASCIICODE(tt.args.base, tt.args.shift); got != tt.want {
				t.Errorf("shiftASCIICODE() = %v, want %v", got, tt.want)
			}
		})
	}
}
