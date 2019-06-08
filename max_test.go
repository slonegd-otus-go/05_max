package max_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	max "github.com/slonegd-otus-go/05_max"
)

type nums struct {
	n int
}

func TestFind(t *testing.T) {

	sliceInt := []interface{}{1, 5, 3}
	sliceFloat := []interface{}{5.3, 1.5, 3.2}
	type man struct {
		name string
		age  int
	}
	sliceStruct := []interface{}{
		man{"Denis", 36},
		man{"Alex", 28},
		man{"Liza", 32},
	}

	tests := []struct {
		name    string
		slice   []interface{}
		greater func(i, j int) bool
		want    interface{}
	}{
		{
			name:  "ints",
			slice: sliceInt,
			greater: func(i, j int) bool {
				return sliceInt[i].(int) > sliceInt[j].(int)
			},
			want: 5,
		},
		{
			name:  "floats",
			slice: sliceFloat,
			greater: func(i, j int) bool {
				return sliceFloat[i].(float64) > sliceFloat[j].(float64)
			},
			want: 5.3,
		},
		{
			name:  "structs age",
			slice: sliceStruct,
			greater: func(i, j int) bool {
				return sliceStruct[i].(man).age > sliceStruct[j].(man).age
			},
			want: man{"Denis", 36},
		},
		{
			name:  "structs name",
			slice: sliceStruct,
			greater: func(i, j int) bool {
				return sliceStruct[i].(man).name > sliceStruct[j].(man).name
			},
			want: man{"Liza", 32},
		},
		{
			name:  "empty slice",
			slice: []interface{}{},
			greater: func(i, j int) bool {
				return false
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := max.Find(tt.slice, tt.greater)
			assert.Equal(t, tt.want, got)
		})
	}
}
