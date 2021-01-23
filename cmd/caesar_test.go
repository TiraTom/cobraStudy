package cmd

import (
	"reflect"
	"testing"

	"github.com/spf13/cobra"
)

func Test_caesar(t *testing.T) {
	type args struct {
		cmd  *cobra.Command
		args []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := caesar(tt.args.cmd, tt.args.args); !reflect.DeepEqual(got, tt.want) {
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
		// TODO: Add test cases.
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
