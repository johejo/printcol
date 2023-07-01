package main

import (
	"reflect"
	"testing"
)

func Test_parseCol(t *testing.T) {
	tests := []struct {
		in      string
		want    []int
		wantErr bool
	}{
		{in: "1", want: []int{1}},
		{in: "1..5", want: []int{1, 2, 3, 4, 5}},
		{in: "1..3,7..9", want: []int{1, 2, 3, 7, 8, 9}},
		{in: "1,5", want: []int{1, 5}},
		{in: ",1,5", want: []int{1, 5}},
		{in: "..", wantErr: true},
		{in: "1..", wantErr: true},
		{in: "..5", wantErr: true},
		{in: "1..-1", want: []int{-1, 0, 1}},
		{in: "-1..1", want: []int{-1, 0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got, err := parseCol(tt.in)
			if (err == nil) && tt.wantErr {
				t.Errorf("should return error")
			}
			if (err != nil) && !tt.wantErr {
				t.Errorf("should not return error %s", err)
			}
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("want=%v, got=%v", tt.want, got)
			}
		})
	}
}
