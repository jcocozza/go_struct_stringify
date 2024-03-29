package gostructstringify

import (
	"math"
	"testing"
)

type testStruct struct {
	A string
	B int
}
type SuperStruct struct {
	E string
	F *testStruct
}
type SuperStructList struct {
	G string
	Lst []*testStruct
}
type Unit string
type MiscFloat float64
func Test_generateStructInstanceCode(t *testing.T) {
	type args struct {
		instance interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"valid", args{&testStruct{ A: "a", B: 1}}, "&gostructstringify.testStruct{A: \"a\", B: 1}"},
		{"valid sub", args{&SuperStruct{ E: "asdf", F:&testStruct{A: "a", B: 1}}}, "&gostructstringify.SuperStruct{E: \"asdf\", F: &gostructstringify.testStruct{A: \"a\", B: 1}}"},
		{"valid sub", args{&SuperStructList{ G: "asdf", Lst: []*testStruct{{A: "a", B: 1}, {A: "a", B: 1}}}}, "&gostructstringify.SuperStructList{G: \"asdf\", Lst: []*gostructstringify.testStruct{&gostructstringify.testStruct{A: \"a\", B: 1}, &gostructstringify.testStruct{A: \"a\", B: 1}}}"},
		{"valid", args{Unit("asdf")}, "gostructstringify.Unit(\"asdf\")"},
		{"valid", args{MiscFloat(1)}, "gostructstringify.MiscFloat(1)"},
	    {"valid", args{math.NaN()}, "math.NaN()"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StructStringify(tt.args.instance); got != tt.want {
				t.Errorf("generateStructInstanceCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
