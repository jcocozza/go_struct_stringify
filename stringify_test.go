package gostructstringify

import (
	"math"
	"testing"
)

type TestUnexported struct {
	A int
	b int
}

type TestStruct struct {
	A string
	B int
}

type TestStruct2 struct {
	A float64
	B float32
}

type SuperStruct struct {
	E string
	F *TestStruct
}
type SuperStructList struct {
	G   string
	Lst []*TestStruct
}
type Unit string
type MiscFloat float64
type MiscInt int

var NilTest *TestUnexported = nil

func Test_generateStructInstanceCode(t *testing.T) {
	type args struct {
		instance interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"nil pointer", args{NilTest}, "(*gostructstringify.TestUnexported)(nil)"},
		{"valid", args{&TestStruct{A: "a", B: 1}}, "&gostructstringify.TestStruct{A: \"a\", B: 1}"},
		{"valid", args{&TestStruct2{A: 1.01, B: 2.5}}, "&gostructstringify.TestStruct2{A: 1.01, B: 2.5}"},
		{"unexported", args{TestUnexported{A: 1, b: 1}}, "gostructstringify.TestUnexported{A: 1, b: 1}"},
		{"unexported_2", args{&TestUnexported{A: 1, b: 1}}, "&gostructstringify.TestUnexported{A: 1, b: 1}"},
		{"valid sub", args{&SuperStruct{E: "asdf", F: &TestStruct{A: "a", B: 1}}}, "&gostructstringify.SuperStruct{E: \"asdf\", F: &gostructstringify.TestStruct{A: \"a\", B: 1}}"},
		{"valid sub", args{&SuperStructList{G: "asdf", Lst: []*TestStruct{{A: "a", B: 1}, {A: "a", B: 1}}}}, "&gostructstringify.SuperStructList{G: \"asdf\", Lst: []*gostructstringify.TestStruct{&gostructstringify.TestStruct{A: \"a\", B: 1}, &gostructstringify.TestStruct{A: \"a\", B: 1}}}"},
		{"valid", args{Unit("asdf")}, "gostructstringify.Unit(\"asdf\")"},
		{"valid", args{MiscFloat(1)}, "gostructstringify.MiscFloat(1)"},
		{"valid", args{MiscInt(1)}, "gostructstringify.MiscInt(1)"},
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
