package treesort_test

import(
	"math/rand"
	"sort"
	"testing"
	"gopl/ch4/treesort"
)

func TestSort(t *testing.T){
	data := make([]int, 50)
	for i := range data{
		data[i] = rand.Int() % 50
	}
	treesort
}