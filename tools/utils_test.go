package tools_test

import (
	"fmt"
	"testing"

	"github.com/nextunit-io/go-tools/tools"
	"github.com/stretchr/testify/assert"
)

type testRemoveElementFromSliceCases[T any] struct {
	Slice         []T
	RemoveObj     T
	ExpectedSlice []T
}

func TestRemoveElementFromSlice(t *testing.T) {
	cases := []testRemoveElementFromSliceCases[any]{
		{
			Slice:         []any{1, 2, 3, 45, 6, 7, 8, 89},
			RemoveObj:     89,
			ExpectedSlice: []any{1, 2, 3, 45, 6, 7, 8},
		},
		{
			Slice:         []any{"test", "test1", "test2", "test3", "test4"},
			RemoveObj:     "test2",
			ExpectedSlice: []any{"test", "test1", "test3", "test4"},
		},
		{
			Slice:         []any{true, true, false, true, false, false},
			RemoveObj:     false,
			ExpectedSlice: []any{true, true, true, false, false},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Run case %d", i), func(t *testing.T) {
			assert.Equal(t, c.ExpectedSlice, tools.RemoveElementFromSlice(c.Slice, c.RemoveObj))
		})
	}
}
