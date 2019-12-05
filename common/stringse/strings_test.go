package stringse_test

import (
	"rank-archive/common/stringse"
	"reflect"
	"testing"
)

func TestDiff(t *testing.T) {
	a := []string{"a", "b", "c", "e"}
	b := []string{"b", "e"}

	expectedResult := []string{"a", "c"}
	realResult := stringse.Diff(a, b)

	if !reflect.DeepEqual(expectedResult, realResult) {
		t.Fatalf("期待结果%s,实际解出%s", expectedResult, realResult)
	}
}
