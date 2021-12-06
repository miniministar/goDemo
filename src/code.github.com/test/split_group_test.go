package test

import (
	"reflect"
	"testing"
)

func TestGroupTest(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}

	testGroup := []testCase{
		{"babcbef", "b", []string{"", "a", "c", "ef"}},
		{"a:b:c", ":", []string{"a", "b", "c"}},
		{"abcef", "bc", []string{"a", "ef"}},
	}

	for _, tc := range testGroup {
		got := Split(tc.str, tc.sep)
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("want:%#v got:%#v\n", tc.want, got)
		}
	}
}

func TestGroupMap(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}

	testGroup := map[string]testCase{
		"case 1": {"babcbef", "b", []string{"", "a", "c", "ef"}},
		"case 2": {"a:b:c", ":", []string{"a", "b", "c"}},
		"case 3": {"abcef", "bc", []string{"a", "ef"}},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.str, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v got:%#v\n", tc.want, got)
			}
		})
	}
}
