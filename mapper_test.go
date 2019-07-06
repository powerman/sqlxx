package sqlxx

import (
	"testing"

	"github.com/powerman/check"
)

func TestToSnake(tt *testing.T) {
	t := check.T(tt)
	t.Parallel()

	testCases := []struct {
		camel string
		snake string
	}{
		{"A", "a"},
		{"AB", "ab"},
		{"ABC", "abc"},
		{"ABCd", "ab_cd"},
		{"ABCdE", "ab_cd_e"},
		{"ABCde", "ab_cde"},
		{"Ab", "ab"},
		{"Abc", "abc"},
		{"AbC", "ab_c"},
		{"AbCd", "ab_cd"},
		{"AbCD", "ab_cd"},
		{"SomeIDOfEntity", "some_id_of_entity"},
	}
	for _, tc := range testCases {
		t.Equal(ToSnake(tc.camel), tc.snake, tc.camel)
	}
}
