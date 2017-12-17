package sort

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuicksortInterfaceToInt(t *testing.T) {
	// TODO
}

func TestQuicksortFindMediaOfThree(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(2, findMedianOfThree(1, 2, 3))
	assert.Equal(2, findMedianOfThree(2, 1, 3))
	assert.Equal(2, findMedianOfThree(1, 3, 2))
}

func TestQuicksortPickPivot(t *testing.T) {
	assert := assert.New(t)

	list := []interface{}{1, 2, 3, 4, 5, 6, 7}
	assert.Equal(0, pickPivot(list, QUICKSORT_PIVOT_FIRST))
	assert.Equal(3, pickPivot(list, QUICKSORT_PIVOT_MIDDLE))
	assert.Equal(6, pickPivot(list, QUICKSORT_PIVOT_LAST))

	list = []interface{}{1, 2, 3, 4, 5, 6, 7, 8}
	assert.Equal(3, pickPivot(list, QUICKSORT_PIVOT_MIDDLE))

	list = []interface{}{1, 2, 3}
	assert.Equal(1, pickPivot(list, QUICKSORT_PIVOT_MEDIAN_OF_THREE))
	list = []interface{}{1, 3, 2}
	assert.Equal(2, pickPivot(list, QUICKSORT_PIVOT_MEDIAN_OF_THREE))
	list = []interface{}{2, 1, 3}
	assert.Equal(0, pickPivot(list, QUICKSORT_PIVOT_MEDIAN_OF_THREE))
	list = []interface{}{"a", "b", "c"}
	assert.Equal(-1, pickPivot(list, QUICKSORT_PIVOT_MEDIAN_OF_THREE))

}

func TestQuicksortWithOptions(t *testing.T) {
	assert := assert.New(t)

	intComparator := func(a, b interface{}) (int, error) {
		valA, ok := a.(int)
		if !ok {
			return 0, errors.New("Compare value must be int")
		}
		valB, ok := b.(int)
		if !ok {
			return 0, errors.New("Compare value must be int")
		}

		if valA > valB {
			return 1, nil
		} else if valA < valB {
			return -1, nil
		} else {
			return 0, nil
		}
	}

	result, err := QuicksortWithOptions(
		[]interface{}{5, 1, 2, 6, 7, 3, 4, 8},
		intComparator,
		QUICKSORT_DEFAULT,
	)
	assert.Equal([]interface{}{1, 2, 3, 4, 5, 6, 7, 8}, result)
	assert.Nil(err)

	result, err = QuicksortWithOptions(
		[]interface{}{8, 7, 6, 5, 4, 3, 2, 1},
		intComparator,
		QUICKSORT_DEFAULT,
	)
	assert.Equal([]interface{}{1, 2, 3, 4, 5, 6, 7, 8}, result)
	assert.Nil(err)

	result, err = QuicksortWithOptions(
		[]interface{}{5, 1, 2, 6, "a", 3, 4, 8},
		intComparator,
		QUICKSORT_DEFAULT,
	)
	assert.Equal("Compare value must be int", err.Error())

	// TODO: Test median-of-three on string array
}
