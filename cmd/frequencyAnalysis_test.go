package cmd

import (
	"reflect"
	"testing"
)

func Test_countChars(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{name: "単一文字の場合", args: args{s: "AAA"}, want: map[string]int{"A": 3}},
		{name: "複数種類の文字の場合", args: args{s: "AB"}, want: map[string]int{"A": 1, "B": 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countChars(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countChars() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_keySort(t *testing.T) {
	type args struct {
		m map[string]int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "並べ替え不要の場合", args: args{m: map[string]int{"a": 2, "b": 3}}, want: []string{"a", "b"}},
		{name: "並べ替えが必要な場合1", args: args{m: map[string]int{"b": 2, "a": 3}}, want: []string{"a", "b"}},
		{name: "並べ替えが必要な場合2", args: args{m: map[string]int{"b": 2, "a": 3, "c": 3}}, want: []string{"a", "b", "c"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := keySort(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("keySort() = %v, want %v", got, tt.want)
			}
		})
	}
}
