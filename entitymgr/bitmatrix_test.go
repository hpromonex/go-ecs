package entitymgr_test

import (
	"testing"

	"github.com/hpromonex/go-ecs/entitymgr"

	"github.com/hpromonex/go-ecs/component"
)

type QueryResult struct {
	input    []component.ID
	expected []uint32
}

var queryTest = func() entitymgr.Bitmatrix {
	var bm entitymgr.Bitmatrix

	//Entity 0 -> A B C
	bm.Set(0, component.A, true)
	bm.Set(0, component.B, true)
	bm.Set(0, component.C, true)

	//Entity 1 -> C D E
	bm.Set(1, component.C, true)
	bm.Set(1, component.D, true)
	bm.Set(1, component.E, true)

	//Entity 2 -> A C F
	bm.Set(2, component.A, true)
	bm.Set(2, component.C, true)
	bm.Set(2, component.F, true)

	//Entity 3 -> F
	bm.Set(3, component.F, true)

	//Entity 4 -> E L
	bm.Set(4, component.L, true)
	bm.Set(4, component.E, true)

	return bm
}()

var queryResults = []QueryResult{
	{
		input:    []component.ID{component.A},
		expected: []uint32{0, 2},
	},
	{
		input:    []component.ID{component.C},
		expected: []uint32{0, 1, 2},
	},
	{
		input:    []component.ID{component.A, component.C},
		expected: []uint32{0, 2},
	},
	{
		input:    []component.ID{component.F},
		expected: []uint32{2, 3},
	},
	{
		input:    []component.ID{component.L},
		expected: []uint32{4},
	},
	{
		input:    []component.ID{component.H},
		expected: []uint32{},
	},
}

func TestBitmatrixQuery(t *testing.T) {

	for _, test := range queryResults {

		result := queryTest.Query(test.input...)

		if len(result) != len(test.expected) {
			t.Errorf("Result '%v' from Input '%v' doesn't match expected '%v'", result, test.input, test.expected)
			continue
		}

		l := len(result)
		if len(test.expected) < l {
			l = len(test.expected)
		}

		for i := 0; i < l; i++ {

			if result[i] != test.expected[i] {
				t.Errorf("Result '%v' from Input '%v' doesn't match expected '%v'", result, test.input, test.expected)
				continue
			}

		}

	}

}
