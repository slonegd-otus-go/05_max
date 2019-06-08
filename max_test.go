package max_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/slonegd-otus-go/05_max"
)

type nums struct {
	n int
}

func TestFind(t *testing.T) {

	sliceInt := []interface{}{1,5,3}
	sliceFloat := []interface{}{5.3,1.5,3.2}
	type man struct {
		name string
		age int
	}
	sliceStruct := []interface{}{
		man{"Denis",36},
		man{"Alex",28},
		man{"Liza",32},
	}

	tests := []struct {
		name string
		slice      []interface{}
		comparator func(i, j int) bool
		want interface{}
	}{
		{
			name : "ints",
			slice : sliceInt,
			comparator: func(i, j int) bool {
				return sliceInt[i].(int) < sliceInt[j].(int)
			},
			want : 5,
		},
		{
			name : "floats",
			slice : sliceFloat,
			comparator: func(i, j int) bool {
				return sliceInt[i].(float64) < sliceInt[j].(float64)
			},
			want : 5.3,
		},
		{
			name : "structs age",
			slice : sliceStruct,
			comparator: func(i, j int) bool {
				return sliceInt[i].(man).age < sliceInt[j].(man).age
			},
			want : man{"Denis",36},
		},
		{
			name : "structs name",
			slice : sliceStruct,
			comparator: func(i, j int) bool {
				return sliceInt[i].(man).name < sliceInt[j].(man).name
			},
			want : man{"Liza",32},
		},
		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := max.Find(tt.slice, tt.comparator)
			assert.Equal(t, tt.want, got)
		})
	}
}
